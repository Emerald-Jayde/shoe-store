package event

import (
	"errors"
	"shoe-store-server/entity"
	"shoe-store-server/lib/pusher"
	"shoe-store-server/repository/sqlite"
	"time"
)

var (
	ErrStoreNotFound     = errors.New("store not found")
	ErrShoeModelNotFound = errors.New("shoe model not found")
)

func CreateSaleEvent(storeName string, shoeModelName string, newAmount int) error {
	store := entity.Store{Name: storeName}
	if err := sqlite.GetStoreByName(&store); err != nil {
		return ErrStoreNotFound
	}

	shoe := entity.ShoeModel{Name: shoeModelName}
	if err := sqlite.GetShoeModelByName(&shoe); err != nil {
		return ErrShoeModelNotFound
	}

	inventory := entity.Inventory{
		StoreID:     store.ID,
		ShoeModelID: shoe.ID,
	}
	sqlite.GetInventory(&inventory)

	sale := createNewSale(inventory, newAmount)
	pusher.PushNewSale(storeName, shoeModelName, sale.NewInventory, sale.OldInventory, sale.CreatedAt)

	updateInventory(&inventory, newAmount)
	pusher.PushInventoryUpdate(int(inventory.ID), inventory.Amount, inventory.UpdatedAt)
	return nil
}

func updateInventory(inventory *entity.Inventory, newAmount int) {
	inventory.Amount = newAmount
	sqlite.UpdateInventory(inventory)
}

func createNewSale(inventory entity.Inventory, newAmount int) entity.Sale {
	sale := entity.Sale{
		StoreID:      inventory.StoreID,
		ShoeModelID:  inventory.ShoeModelID,
		NewInventory: newAmount,
		OldInventory: inventory.Amount,
		CreatedAt:    time.Now(),
	}
	sqlite.CreateSale(&sale)

	return sale
}
