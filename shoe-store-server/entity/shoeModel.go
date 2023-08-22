package entity

import (
	"gorm.io/gorm"
)

type ShoeModel struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:idx_shoe_model_name"`
}
