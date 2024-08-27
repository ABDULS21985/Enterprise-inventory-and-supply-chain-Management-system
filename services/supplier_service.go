package services

import (
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
)

func CreateSupplier(supplier models.Supplier) error {
	return db.DB.Create(&supplier).Error
}

func GetAllSuppliers() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	err := db.DB.Find(&suppliers).Error
	return suppliers, err
}
