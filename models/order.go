package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      uint    `json:"user_id"`
	InventoryID uint    `json:"inventory_id"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
	Status      string  `json:"status"`
}
