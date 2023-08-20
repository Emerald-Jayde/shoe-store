package entity

import (
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:idx_store_name"`
}
