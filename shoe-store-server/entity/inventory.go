package entity

import (
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	StoreID     int
	Store       Store
	ShoeModelID int
	ShoeModel   ShoeModel
	Amount      int
}
