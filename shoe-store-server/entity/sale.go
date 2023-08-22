package entity

import "time"

type Sale struct {
	StoreID      uint `gorm:"index:idx_sales_store_id"`
	ShoeModelID  uint `gorm:"index:idx_sales_shoe_model_id"`
	NewInventory int
	OldInventory int
	CreatedAt    time.Time
}
