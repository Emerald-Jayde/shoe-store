package sqlite

import (
	"shoe-store-server/db"
	"shoe-store-server/entity"
)

func CreateShoeModel(shoe *entity.ShoeModel) {
	db.GetDBInstance().Create(shoe)
}
func GetShoeModel(shoe *entity.ShoeModel) error {
	tx := db.GetDBInstance().First(shoe)
	return tx.Error
}

func GetShoeModels(shoes *[]entity.ShoeModel) {
	db.GetDBInstance().Find(shoes)
}

func GetShoeModelNameById(shoeModelId uint) string {
	sm := entity.ShoeModel{}
	db.GetDBInstance().
		Model(&entity.ShoeModel{}).
		Select("name").
		First(&sm, shoeModelId)
	return sm.Name
}

func GetShoeModelByName(shoe *entity.ShoeModel) error {
	tx := db.GetDBInstance().Where("name=?", shoe.Name).First(shoe)
	return tx.Error
}
