package services

import (
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
)

func CreateInventoryItem(inventory models.Inventory) error {
	result := db.DB.Create(&inventory)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
