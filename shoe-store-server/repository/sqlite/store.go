package sqlite

import (
	"gorm.io/gorm"
	"shoe-store-server/db"
	"shoe-store-server/entity"
)

func CreateStore(store *entity.Store) {
	db.GetDBInstance().Create(store)
}

func GetStore(store *entity.Store) error {
	tx := db.GetDBInstance().First(store)
	return tx.Error
}

func GetStores(stores *[]entity.Store) *gorm.DB {
	return db.GetDBInstance().Find(stores)
}

func GetStoreNameById(storeId uint) string {
	st := entity.Store{}
	db.GetDBInstance().
		Model(&entity.Store{}).
		Select("name").
		First(&st, storeId)
	return st.Name
}

func GetStoreByName(store *entity.Store) error {
	tx := db.GetDBInstance().Where("name = ?", store.Name).First(store)
	return tx.Error
}
