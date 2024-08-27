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

// GetItemByID retrieves an item by ID
func GetItemByID(id uint) (models.Item, error) {
	var item models.Item
	err := db.DB.Preload("Supplier").First(&item, id).Error
	return item, err
}

// UpdateItem updates an existing item
func UpdateItem(item models.Item) error {
	return db.DB.Save(&item).Error
}

// DeleteItem deletes an item by ID
func DeleteItem(id uint) error {
	return db.DB.Delete(&models.Item{}, id).Error
}
