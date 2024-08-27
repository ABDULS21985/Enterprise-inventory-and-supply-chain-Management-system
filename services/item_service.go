package services

import (
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
)

func CreateItem(item models.Item) error {
	return db.DB.Create(&item).Error
}

func GetAllItems() ([]models.Item, error) {
	var items []models.Item
	err := db.DB.Preload("Supplier").Find(&items).Error
	return items, err
}
