package sqlite

import (
	"gorm.io/gorm"
	"shoe-store-server/db"
	"shoe-store-server/entity"
)

func CreateSale(sale *entity.Sale) {
	db.GetDBInstance().Create(sale)
}

func GetSales(sales *[]entity.Sale) {
	db.GetDBInstance().Find(sales)
}

func GetLatestSales(sales *[]entity.Sale, limit int) {
	db.GetDBInstance().Limit(limit).Find(sales)
}

func GetSalesByStoreId(sales *[]entity.Sale, storeId uint) {
	db.GetDBInstance().Find(sales, "store_id=?", storeId)
}

func GetNumberOfSalesByShoeModelId(count *int64, shoeModelId uint) *gorm.DB {
	return db.GetDBInstance().Model(&entity.Sale{}).Where("shoe_model_id=?", shoeModelId).Count(count)
}

func GetNumberOfSalesByStoreId(count *int64, storeId uint) *gorm.DB {
	return db.GetDBInstance().Model(&entity.Sale{}).Where("store_id=?", storeId).Count(count)
}

func GetBestAndWorstShoeModelSalesByStoreId(storeId uint) map[string][]int {
	type ShoeModelSales struct {
		ShoeModelID int
		NSales      int
	}
	s := []ShoeModelSales{}
	db.GetDBInstance().
		Model(&entity.Sale{}).
		Select("COUNT(*) as n_sales, shoe_model_id").
		Where("store_id=?", storeId).
		Group("shoe_model_id").
		Order("n_sales").
		Scan(&s)

	return map[string][]int{
		"worst": {s[0].ShoeModelID, s[0].NSales},
		"best":  {s[len(s)-1].ShoeModelID, s[len(s)-1].NSales},
	}
}
