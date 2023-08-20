package sqlite

import (
	"shoe-store-server/db"
	"shoe-store-server/entity"
)

func CreateSale(sale *entity.Sale) {
	db.GetDBInstance().Create(sale)
}

func GetSales(sales *[]entity.Sale) {
	db.GetDBInstance().Find(sales)
}

func GetSalesByStoreId(sales *[]entity.Sale, storeId uint) {
	db.GetDBInstance().Find(sales, "store_id=?", storeId)
}
