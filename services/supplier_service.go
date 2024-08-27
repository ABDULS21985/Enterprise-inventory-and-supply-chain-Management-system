package services

import (
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
)

// CreateSupplier creates a new supplier
func CreateSupplier(supplier models.Supplier) error {
	result := db.DB.Create(&supplier)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetSuppliers fetches all suppliers from the database
func GetSuppliers() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}

	return suppliers, nil
}

// GetSupplierByID fetches a supplier by its ID
func GetSupplierByID(id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	result := db.DB.First(&supplier, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &supplier, nil
}

// UpdateSupplier updates an existing supplier in the database
func UpdateSupplier(supplier *models.Supplier) error {
	result := db.DB.Save(supplier)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteSupplier deletes a supplier from the database by its ID
func DeleteSupplier(id uint) error {
	result := db.DB.Delete(&models.Supplier{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetSuppliersByName fetches suppliers by name (supports partial matching)
func GetSuppliersByName(name string) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("name LIKE ?", "%"+name+"%").Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}

	return suppliers, nil
}

// GetSuppliersByLocation fetches suppliers by location
func GetSuppliersByLocation(location string) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Where("location = ?", location).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}

	return suppliers, nil
}

// GetSuppliersByProductID fetches suppliers that provide a specific product by its product ID
func GetSuppliersByProductID(productID uint) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := db.DB.Joins("JOIN products ON suppliers.id = products.supplier_id").Where("products.id = ?", productID).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}

	return suppliers, nil
}
