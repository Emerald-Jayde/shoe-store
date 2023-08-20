package sqlite

import (
	"shoe-store-server/db"
	"shoe-store-server/entity"
)

func CreateInventory(inventory *entity.Inventory) {
	db.GetDBInstance().Create(inventory)
}

func GetInventoriesByStoreId(inventories *[]entity.Inventory, storeId uint) {
	db.GetDBInstance().
		Preload("Store").
		Preload("ShoeModel").
		Find(inventories, "store_id = ?", storeId)
}

func GetInventory(inventory *entity.Inventory) {
	db.GetDBInstance().
		Preload("Store").
		Preload("ShoeModel").
		First(inventory, "store_id=? AND shoe_model_id=?", inventory.StoreID, inventory.ShoeModelID)
}

func GetInventories(inventories *[]entity.Inventory) {
	db.GetDBInstance().
		Preload("Store").
		Preload("ShoeModel").
		Find(inventories)
}

func UpdateInventory(inventory *entity.Inventory) {
	db.GetDBInstance().Save(inventory)
}
