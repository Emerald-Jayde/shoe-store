package pusher

import (
	"fmt"
	"github.com/pusher/pusher-http-go/v5"
	"time"
)

type PusherSale struct {
	Store     string    `json:"store"`
	ShoeModel string    `json:"shoe_model"`
	NewAmount int       `json:"new_amount"`
	OldAmount int       `json:"old_amount"`
	CreatedAt time.Time `json:"created_at"`
}

type PusherInventory struct {
	InventoryID int       `json:"inventory_id"`
	NewAmount   int       `json:"new_amount"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var PusherClient *pusher.Client

func SetupPusher() {
	pusherClient := pusher.Client{
		AppID:   "1655017",
		Key:     "f8fae05c56b80676064a",
		Secret:  "f5017f13d28a9f235e4a",
		Cluster: "us2",
		Secure:  true,
	}

	PusherClient = &pusherClient
}

func PushNewSale(store string, shoeModel string, newAmount int, oldAmount int, createdAt time.Time) {
	data := PusherSale{
		Store:     store,
		ShoeModel: shoeModel,
		NewAmount: newAmount,
		OldAmount: oldAmount,
		CreatedAt: createdAt,
	}

	if err := PusherClient.Trigger("shoe-store-potloc", "newSale", data); err != nil {
		fmt.Println(err.Error())
	}
}

func PushInventoryUpdate(inventoryId int, newAmount int, updatedAt time.Time) {
	data := PusherInventory{
		InventoryID: inventoryId,
		NewAmount:   newAmount,
		UpdatedAt:   updatedAt,
	}

	if err := PusherClient.Trigger("shoe-store-potloc", "inventoryUpdate", data); err != nil {
		fmt.Println(err.Error())
	}
}
