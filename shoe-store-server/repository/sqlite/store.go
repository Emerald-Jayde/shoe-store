package sqlite

import (
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

func GetStores(stores *[]entity.Store) {
	db.GetDBInstance().Find(stores)
}

func GetStoreName(store *entity.Store) error {
	tx := db.GetDBInstance().Select("name").First(store)
	return tx.Error
}

func GetStoreByName(store *entity.Store) error {
	tx := db.GetDBInstance().Where("name = ?", store.Name).First(store)
	return tx.Error
}
