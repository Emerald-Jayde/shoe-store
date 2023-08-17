package entity

import "gorm.io/gorm"

type ShoeModel struct {
	gorm.Model
	Name string
}
