package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"net/url"
	"shoe-store-server/event"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Client return websocket client connection
type Client struct {
	configStr string
	ctx       context.Context
	ctxCancel context.CancelFunc

	mu     sync.RWMutex
	wsconn *websocket.Conn
}

type message struct {
	Store     string `json:"store"`
	Model     string `json:"model"`
	Inventory int    `json:"inventory"`
}

func StartWebsocket(host, channel string) error {
	client, err := NewWebsocketClient(host, channel)
	client.Connect()

	return err
}

// NewWebsocketClient create new websocket connection
func NewWebsocketClient(host, channel string) (*Client, error) {
	conn := Client{}
	conn.ctx, conn.ctxCancel = context.WithCancel(context.Background())

	u := url.URL{Scheme: "ws", Host: host, Path: channel}
	conn.configStr = u.String()

	go conn.listen()
	return &conn, nil
}

func (conn *Client) Connect() *websocket.Conn {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	if conn.wsconn != nil {
		return conn.wsconn
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		select {
		case <-conn.ctx.Done():
			return nil
		default:
			ws, _, err := websocket.DefaultDialer.Dial(conn.configStr, nil)
			if err != nil {
				conn.log("connect", err, fmt.Sprintf("Cannot connect to websocket: %s", conn.configStr))
				continue
			}
			conn.log("connect", nil, fmt.Sprintf("connected to websocket to %s", conn.configStr))
			conn.wsconn = ws
			return conn.wsconn
		}
	}
}

func (conn *Client) listen() {
	conn.log("listen", nil, fmt.Sprintf("listen for the messages: %s", conn.configStr))
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-conn.ctx.Done():
			return
		case <-ticker.C:
			for {
				ws := conn.Connect()
				if ws == nil {
					return
				}
				_, wsMsg, readErr := ws.ReadMessage()
				if readErr != nil {
					conn.log("listen", readErr, "Cannot read websocket message")
					conn.closeWs()
					break
				}
				conn.log("listen", nil, fmt.Sprintf("websocket msg: %s\n", string(wsMsg)))

				conn.createNewSaleEvent(wsMsg)
			}
		}
	}
}

// Stop will send close message and shutdown websocket connection
func (conn *Client) Stop() {
	conn.ctxCancel()
	conn.closeWs()
}

// closeWs will send close message and shutdown websocket connection
func (conn *Client) closeWs() {
	conn.mu.Lock()
	if conn.wsconn != nil {
		_ = conn.wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		_ = conn.wsconn.Close()
		conn.wsconn = nil
	}
	conn.mu.Unlock()
	conn.log("closeWs", nil, "closed websocket")
}

func (conn *Client) createNewSaleEvent(wsMessage []byte) {
	// create new sale event
	var m message
	json.Unmarshal(wsMessage, &m)
	err := event.CreateSaleEvent(m.Store, m.Model, m.Inventory)
	conn.log("createNewSaleEvent", err, "Creating SaleEvent...")
}

func (conn *Client) log(funcName string, err error, msg string) {
	if err != nil {
		log.Errorf("[WS Client Error] func %s: err: %v, msg: %s\n", funcName, err, msg)
	} else {
		log.Infof("[WS Client] func %s: %s\n", funcName, msg)
	}
}
