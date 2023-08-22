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

func GetHighStockInventory(i *entity.Inventory) {
	db.GetDBInstance().
		Preload("Store").
		Select("store_id, max(amount) as amount").
		Where("shoe_model_id=? and amount > 50", i.ShoeModelID).
		Order("shoe_model_id DESC").
		First(&i)
}

func GetLowStockInventories(i *[]entity.Inventory) {
	db.GetDBInstance().
		Preload("Store").Preload("ShoeModel").Raw(`
			SELECT store_id, shoe_model_id, amount
			FROM (
					 SELECT *, ROW_NUMBER() OVER (PARTITION BY shoe_model_id ORDER BY amount ASC) AS n
					 FROM inventories
					 WHERE amount < 11
				 ) AS x
			WHERE n = 1
		`).Find(&i)
}

func GetInventoryByStoreAndShoeModel(inventory *entity.Inventory) {
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
