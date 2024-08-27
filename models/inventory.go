package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Name     string  `json:"name"`
	SKU      string  `json:"sku" gorm:"unique"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	VendorID uint    `json:"vendor_id"`
}
