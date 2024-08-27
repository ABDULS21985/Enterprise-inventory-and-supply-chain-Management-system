package services

import (
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
)

// CreateItem adds a new item to the database.
func CreateItem(item models.Item) error {
	result := db.DB.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetItems fetches all items from the database.
func GetItems() ([]models.Item, error) {
	var items []models.Item
	result := db.DB.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemByID fetches a single item by its ID.
func GetItemByID(id uint) (*models.Item, error) {
	var item models.Item
	result := db.DB.First(&item, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

// UpdateItem updates an existing item in the database.
func UpdateItem(item *models.Item) error {
	result := db.DB.Save(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteItem removes an item from the database using its ID.
func DeleteItem(id uint) error {
	result := db.DB.Delete(&models.Item{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetItemsByCategory fetches items based on their category.
func GetItemsByCategory(category string) ([]models.Item, error) {
	var items []models.Item
	result := db.DB.Where("category = ?", category).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemsByWarehouseID fetches items stored in a specific warehouse.
func GetItemsByWarehouseID(warehouseID uint) ([]models.Item, error) {
	var items []models.Item
	result := db.DB.Where("warehouse_id = ?", warehouseID).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemsBySupplierID fetches items supplied by a specific supplier.
func GetItemsBySupplierID(supplierID uint) ([]models.Item, error) {
	var items []models.Item
	result := db.DB.Where("supplier_id = ?", supplierID).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemsByStockRange fetches items within a specified stock level range.
func GetItemsByStockRange(minStock, maxStock int) ([]models.Item, error) {
	var items []models.Item
	result := db.DB.Where("stock >= ? AND stock <= ?", minStock, maxStock).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemsByPriceRange fetches items within a specified price range.
func GetItemsByPriceRange(minPrice, maxPrice float64) ([]models.Item, error) {
	var items []models.Item
	result := db.DB.Where("price >= ? AND price <= ?", minPrice, maxPrice).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemsByCategoryAndPriceRange fetches items based on their category and price range.
func GetItemsByCategoryAndPriceRange(category string, minPrice, maxPrice float64) ([]models.Item, error) {
	var items []models.Item
	result := db.DB.Where("category = ? AND price >= ? AND price <= ?", category, minPrice, maxPrice).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemsByCategoryAndStockRange fetches items based on their category and stock level range.
func GetItemsByCategoryAndStockRange(category string, minStock, maxStock int) ([]models.Item, error) {
	var items []models.Item
	result := db.DB.Where("category = ? AND stock >= ? AND stock <= ?", category, minStock, maxStock).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemsByCategoryAndSupplierID fetches items based on their category and supplier ID.
func GetItemsByCategoryAndSupplierID(category string, supplierID uint) ([]models.Item, error) {

	var items []models.Item
	result := db.DB.Where("category = ? AND supplier_id = ?", category, supplierID).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemsByCategoryAndWarehouseID fetches items based on their category and warehouse ID.
func GetItemsByCategoryAndWarehouseID(category string, warehouseID uint) ([]models.Item, error) {
	var items []models.Item
	result := db.DB.Where("category = ? AND warehouse_id = ?", category, warehouseID).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemsBySupplierIDAndWarehouseID fetches items based on their supplier ID and warehouse ID.

func GetItemsBySupplierIDAndWarehouseID(supplierID, warehouseID uint) ([]models.Item, error) {

	var items []models.Item
	result := db.DB.Where("supplier_id = ? AND warehouse_id = ?", supplierID, warehouseID).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemsBySupplierIDAndPriceRange fetches items based on their supplier ID and price range.

func GetItemsBySupplierIDAndPriceRange(supplierID uint, minPrice, maxPrice float64) ([]models.Item, error) {

	var items []models.Item
	result := db.DB.Where("supplier_id = ? AND price >= ? AND price <= ?", supplierID, minPrice, maxPrice).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemsBySupplierIDAndStockRange fetches items based on their supplier ID and stock level range.

func GetItemsBySupplierIDAndStockRange(supplierID uint, minStock, maxStock int) ([]models.Item, error) {

	var items []models.Item
	result := db.DB.Where("supplier_id = ? AND stock >= ? AND stock <= ?", supplierID, minStock, maxStock).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}
