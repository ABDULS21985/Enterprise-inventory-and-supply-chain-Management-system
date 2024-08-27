package services

import (
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
)

// CreateProduct creates a new product
func CreateProduct(product models.Product) error {
	return db.DB.Create(&product).Error
}

// GetProducts fetches all products from the database
func GetProducts() ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

// GetProductByID fetches a product by its ID
func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	result := db.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

// UpdateProduct updates a product in the database
func UpdateProduct(product *models.Product) error {
	return db.DB.Save(product).Error
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id uint) error {
	return db.DB.Delete(&models.Product{}, id).Error
}

// GetProductsByCategory fetches all products for a given category
func GetProductsByCategory(category string) ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Where("category = ?", category).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

// GetProductsByPriceRange fetches all products within a given price range
func GetProductsByPriceRange(minPrice, maxPrice float64) ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Where("price >= ? AND price <= ?", minPrice, maxPrice).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

// GetProductsByStock fetches all products with a given stock level
func GetProductsByStock(stock int) ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Where("stock = ?", stock).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

// GetProductsByStockRange fetches all products within a given stock range
func GetProductsByStockRange(minStock, maxStock int) ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Where("stock >= ? AND stock <= ?", minStock, maxStock).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

// GetProductsByCategoryAndPriceRange fetches all products for a given category within a given price range
func GetProductsByCategoryAndPriceRange(category string, minPrice, maxPrice float64) ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Where("category = ? AND price >= ? AND price <= ?", category, minPrice, maxPrice).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

// GetProductsByCategoryAndStock fetches all products for a given category with a given stock level
func GetProductsByCategoryAndStock(category string, stock int) ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Where("category = ? AND stock = ?", category, stock).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

// GetProductsByCategoryAndStockRange fetches all products for a given category within a given stock range

func GetProductsByCategoryAndStockRange(category string, minStock, maxStock int) ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Where("category = ? AND stock >= ? AND stock <= ?", category, minStock, maxStock).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

// GetProductsByPriceRangeAndStock fetches all products within a given price range with a given stock level
func GetProductsByPriceRangeAndStock(minPrice, maxPrice float64, stock int) ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Where("price >= ? AND price <= ? AND stock = ?", minPrice, maxPrice, stock).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

// GetProductsByPriceRangeAndStockRange fetches all products within a given price range and stock range
func GetProductsByPriceRangeAndStockRange(minPrice, maxPrice float64, minStock, maxStock int) ([]models.Product, error) {

	var products []models.Product
	result := db.DB.Where("price >= ? AND price <= ? AND stock >= ? AND stock <= ?", minPrice, maxPrice, minStock, maxStock).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

// GetProductsByCategoryPriceRangeAndStock fetches all products for a given category within a given price range with a given stock level
func GetProductsByCategoryPriceRangeAndStock(category string, minPrice, maxPrice float64, stock int) ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Where("category = ? AND price >= ? AND price <= ? AND stock = ?", category, minPrice, maxPrice, stock).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

// GetProductsByCategoryPriceRangeAndStockRange fetches all products for a given category within a given price range and stock range
func GetProductsByCategoryPriceRangeAndStockRange(category string, minPrice, maxPrice float64, minStock, maxStock int) ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Where("category = ? AND price >= ? AND price <= ? AND stock >= ? AND stock <= ?", category, minPrice, maxPrice, minStock, maxStock).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
