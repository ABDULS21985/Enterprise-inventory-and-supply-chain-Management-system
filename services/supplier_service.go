package services

import (
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
)

// CreateSupplier adds a new supplier to the database.
func CreateSupplier(supplier models.Supplier) error {
	result := db.DB.Create(&supplier)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetSuppliers fetches all suppliers from the database.
func GetSuppliers() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSupplierByID fetches a supplier by its ID.
func GetSupplierByID(id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	result := db.DB.First(&supplier, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &supplier, nil
}

// UpdateSupplier updates an existing supplier in the database.
func UpdateSupplier(supplier *models.Supplier) error {
	result := db.DB.Save(supplier)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteSupplier deletes a supplier from the database using its ID.
func DeleteSupplier(id uint) error {
	result := db.DB.Delete(&models.Supplier{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetSuppliersByCategory fetches suppliers based on their category.
func GetSuppliersByCategory(category string) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("category = ?", category).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByProductID fetches suppliers based on the products they supply.
func GetSuppliersByProductID(productID uint) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("product_id = ?", productID).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByLocation fetches suppliers based on their location.
func GetSuppliersByLocation(location string) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("location = ?", location).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByRating fetches suppliers based on their rating.
func GetSuppliersByRating(rating float32) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("rating = ?", rating).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByProductIDAndLocation fetches suppliers based on the products they supply and their location.
func GetSuppliersByProductIDAndLocation(productID uint, location string) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("product_id = ? AND location = ?", productID, location).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByProductIDAndRating fetches suppliers based on the products they supply and their rating.
func GetSuppliersByProductIDAndRating(productID uint, rating float32) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("product_id = ? AND rating = ?", productID, rating).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByLocationAndRating fetches suppliers based on their location and rating.
func GetSuppliersByLocationAndRating(location string, rating float32) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("location = ? AND rating = ?", location, rating).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByProductIDLocationAndRating fetches suppliers based on the products they supply, their location, and rating.
func GetSuppliersByProductIDLocationAndRating(productID uint, location string, rating float32) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("product_id = ? AND location = ? AND rating = ?", productID, location, rating).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByCategoryLocationAndRating fetches suppliers based on their category, location, and rating.
func GetSuppliersByCategoryLocationAndRating(category, location string, rating float32) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("category = ? AND location = ? AND rating = ?", category, location, rating).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByCategoryLocationRatingAndProductID fetches suppliers based on their category, location, rating, and product ID.
func GetSuppliersByCategoryLocationRatingAndProductID(category, location string, rating float32, productID uint) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("category = ? AND location = ? AND rating = ? AND product_id = ?", category, location, rating, productID).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByOrderCount fetches suppliers based on the number of orders they've fulfilled.
func GetSuppliersByOrderCount(minOrders, maxOrders int) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("order_count >= ? AND order_count <= ?", minOrders, maxOrders).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByCategoryAndLocation fetches suppliers based on their category and location.
func GetSuppliersByCategoryAndLocation(category, location string) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("category = ? AND location = ?", category, location).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByCategoryAndRating fetches suppliers based on their category and rating.
func GetSuppliersByCategoryAndRating(category string, minRating, maxRating float64) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("category = ? AND rating >= ? AND rating <= ?", category, minRating, maxRating).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByCategoryAndOrderCount fetches suppliers based on their category and the number of orders they've fulfilled.
func GetSuppliersByCategoryAndOrderCount(category string, minOrders, maxOrders int) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("category = ? AND order_count >= ? AND order_count <= ?", category, minOrders, maxOrders).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByLocationAndOrderCount fetches suppliers based on their location and the number of orders they've fulfilled.
func GetSuppliersByLocationAndOrderCount(location string, minOrders, maxOrders int) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("location = ? AND order_count >= ? AND order_count <= ?", location, minOrders, maxOrders).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByRatingAndOrderCount fetches suppliers based on their rating and the number of orders they've fulfilled.
func GetSuppliersByRatingAndOrderCount(minRating, maxRating float64, minOrders, maxOrders int) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("rating >= ? AND rating <= ? AND order_count >= ? AND order_count <= ?", minRating, maxRating, minOrders, maxOrders).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByProductIDAndLocationAndRating fetches suppliers based on the product ID, location, and rating.
func GetSuppliersByProductIDAndLocationAndRating(productID uint, location string, rating float32) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("product_id = ? AND location = ? AND rating = ?", productID, location, rating).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSuppliersByCategoryAndLocationAndRating fetches suppliers based on the category, location, and rating.
func GetSuppliersByCategoryAndLocationAndRating(category string, location string, rating float32) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("category = ? AND location = ? AND rating = ?", category, location, rating).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}
