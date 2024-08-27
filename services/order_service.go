package services

import (
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
)

func CreateOrder(order models.Order) error {
	result := db.DB.Create(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
