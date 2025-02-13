package pusher

import (
	"fmt"
	"github.com/pusher/pusher-http-go/v5"
	"os"
	"time"
)

type sale struct {
	Store     string    `json:"store"`
	ShoeModel string    `json:"shoe_model"`
	NewAmount int       `json:"new_amount"`
	OldAmount int       `json:"old_amount"`
	CreatedAt time.Time `json:"created_at"`
}

type inventory struct {
	InventoryID int       `json:"inventory_id"`
	NewAmount   int       `json:"new_amount"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const (
	potlocChannel  = "shoe-store-potloc"
	newSaleEvent   = "newSale"
	invUpdateEvent = "inventoryUpdate"
)

var Client *pusher.Client

func SetupPusher() {
	pusherClient := pusher.Client{
		AppID:   os.Getenv("PUSHER_APPID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: "us2",
		Secure:  true,
	}

	Client = &pusherClient
}

func PushNewSale(store string, shoeModel string, newAmount int, oldAmount int, createdAt time.Time) {
	data := sale{
		Store:     store,
		ShoeModel: shoeModel,
		NewAmount: newAmount,
		OldAmount: oldAmount,
		CreatedAt: createdAt,
	}

	if err := Client.Trigger(potlocChannel, newSaleEvent, data); err != nil {
		fmt.Println(err.Error())
	}
}

func PushInventoryUpdate(inventoryId int, newAmount int, updatedAt time.Time) {
	data := inventory{
		InventoryID: inventoryId,
		NewAmount:   newAmount,
		UpdatedAt:   updatedAt,
	}

	if err := Client.Trigger(potlocChannel, invUpdateEvent, data); err != nil {
		fmt.Println(err.Error())
	}
}
