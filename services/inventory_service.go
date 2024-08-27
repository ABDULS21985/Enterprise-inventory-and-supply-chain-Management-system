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

// GetInventoryItems fetches all inventory items from the database
func GetInventoryItems() ([]models.Inventory, error) {
	var inventoryItems []models.Inventory
	result := db.DB.Find(&inventoryItems)
	if result.Error != nil {
		return nil, result.Error
	}

	return inventoryItems, nil
}

// GetInventoryItemByID fetches an inventory item by its ID
func GetInventoryItemByID(id uint) (*models.Inventory, error) {
	var inventoryItem models.Inventory
	result := db.DB.First(&inventoryItem, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &inventoryItem, nil
}

// UpdateInventoryItem updates an inventory item in the database
func UpdateInventoryItem(inventory *models.Inventory) error {
	result := db.DB.Save(inventory)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteInventoryItem deletes an inventory item from the database
func DeleteInventoryItem(id uint) error {
	result := db.DB.Delete(&models.Inventory{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetInventoryItemsByProductID fetches all inventory items for a given product ID
func GetInventoryItemsByProductID(productID uint) ([]models.Inventory, error) {
	var inventoryItems []models.Inventory
	result := db.DB.Where("product_id = ?", productID).Find(&inventoryItems)
	if result.Error != nil {
		return nil, result.Error
	}

	return inventoryItems, nil
}

// GetInventoryItemsByWarehouseID fetches all inventory items for a given warehouse ID
func GetInventoryItemsByWarehouseID(warehouseID uint) ([]models.Inventory, error) {
	var inventoryItems []models.Inventory
	result := db.DB.Where("warehouse_id = ?", warehouseID).Find(&inventoryItems)
	if result.Error != nil {
		return nil, result.Error
	}

	return inventoryItems, nil
}
