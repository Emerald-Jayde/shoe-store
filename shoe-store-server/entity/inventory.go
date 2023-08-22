package entity

import (
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	StoreID     uint `gorm:"index:idx_inventory_store_id,index:idx_store_and_shoe_ids"`
	Store       Store
	ShoeModelID uint `gorm:"index:idx_inventory_shoe_model_id,index:idx_store_and_shoe_ids"`
	ShoeModel   ShoeModel
	Amount      int
}
