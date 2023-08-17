package initializers

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"shoe-store-server/helpers"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gorilla/websocket"
)

// TODO: change the name to Sale or InventoryResponse
type Response struct {
	Store     string `json:"store"`
	Model     string `json:"model"`
	Inventory int    `json:"inventory"`
}

func WebsocketClient() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	wsURL := os.Getenv("WEBSOCKET_URL")

	log.Infof("connecting to %s", wsURL)

	// Establishes connection
	c, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	helpers.HandleError("dial: ", err, true)

	defer func(c *websocket.Conn) {
		err := c.Close()
		helpers.HandleError("close websocket connection:", err, true)
	}(c)

	done := make(chan struct{})

	// go routine deals with concurrency
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			helpers.HandleError("read:", err, true)

			log.Infof("Message received: %s", message)
			populateDB(message)
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Info("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			helpers.HandleError("write close: ", err, true)

			// when the interrupt occurs, this allows us to log cleanly that the ws has been closed (from line 38)
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

// TODO: move this to the right directory
func populateDB(message []byte) {
	var r Response
	json.Unmarshal(message, &r)
	fmt.Println(r)
}
