package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name       string   `json:"name"`
	Quantity   int      `json:"quantity"`
	Price      float64  `json:"price"`
	SupplierID uint     `json:"supplier_id"`
	Supplier   Supplier `gorm:"foreignKey:SupplierID"`
}
