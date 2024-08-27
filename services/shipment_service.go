package services

import (
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
)

// CreateShipment creates a new shipment
func CreateShipment(shipment models.Shipment) error {
	result := db.DB.Create(&shipment)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetShipments fetches all shipments from the database
func GetShipments() ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentByID fetches a shipment by its ID
func GetShipmentByID(id uint) (*models.Shipment, error) {
	var shipment models.Shipment
	result := db.DB.First(&shipment, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &shipment, nil
}

// UpdateShipment updates a shipment in the database
func UpdateShipment(shipment *models.Shipment) error {
	result := db.DB.Save(shipment)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteShipment deletes a shipment from the database
func DeleteShipment(id uint) error {
	result := db.DB.Delete(&models.Shipment{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetShipmentsByStatus fetches all shipments for a given status
func GetShipmentsByStatus(status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("status = ?", status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByProductID fetches all shipments for a given product ID
func GetShipmentsByProductID(productID uint) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("product_id = ?", productID).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByDestination fetches all shipments for a given destination
func GetShipmentsByDestination(destination string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("destination = ?", destination).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByOrigin fetches all shipments for a given origin
func GetShipmentsByOrigin(origin string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("origin = ?", origin).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByWarehouseID fetches all shipments for a given warehouse ID
func GetShipmentsByWarehouseID(warehouseID uint) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("warehouse_id = ?", warehouseID).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByCarrier fetches all shipments for a given carrier
func GetShipmentsByCarrier(carrier string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("carrier = ?", carrier).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByTrackingNumber fetches a shipment by its tracking number
func GetShipmentsByTrackingNumber(trackingNumber string) (*models.Shipment, error) {
	var shipment models.Shipment
	result := db.DB.Where("tracking_number = ?", trackingNumber).First(&shipment)
	if result.Error != nil {
		return nil, result.Error
	}

	return &shipment, nil
}

// GetShipmentsByWarehouseIDAndStatus fetches all shipments for a given warehouse ID and status
func GetShipmentsByWarehouseIDAndStatus(warehouseID uint, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("warehouse_id = ? AND status = ?", warehouseID, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByProductIDAndStatus fetches all shipments for a given product ID and status
func GetShipmentsByProductIDAndStatus(productID uint, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("product_id = ? AND status = ?", productID, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByCarrierAndStatus fetches all shipments for a given carrier and status
func GetShipmentsByCarrierAndStatus(carrier, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("carrier = ? AND status = ?", carrier, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByOriginAndStatus fetches all shipments for a given origin and status
func GetShipmentsByOriginAndStatus(origin, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("origin = ? AND status = ?", origin, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByDestinationAndStatus fetches all shipments for a given destination and status
func GetShipmentsByDestinationAndStatus(destination, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("destination = ? AND status = ?", destination, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByWarehouseIDAndCarrier fetches all shipments for a given warehouse ID and carrier
func GetShipmentsByWarehouseIDAndCarrier(warehouseID uint, carrier string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("warehouse_id = ? AND carrier = ?", warehouseID, carrier).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByWarehouseIDAndCarrierAndStatus fetches all shipments for a given warehouse ID, carrier, and status

func GetShipmentsByWarehouseIDAndCarrierAndStatus(warehouseID uint, carrier, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("warehouse_id = ? AND carrier = ? AND status = ?", warehouseID, carrier, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByWarehouseIDAndOrigin fetches all shipments for a given warehouse ID and origin
func GetShipmentsByWarehouseIDAndOrigin(warehouseID uint, origin string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("warehouse_id = ? AND origin = ?", warehouseID, origin).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByWarehouseIDAndOriginAndStatus fetches all shipments for a given warehouse ID, origin, and status
func GetShipmentsByWarehouseIDAndOriginAndStatus(warehouseID uint, origin, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("warehouse_id = ? AND origin = ? AND status = ?", warehouseID, origin, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByWarehouseIDAndDestination fetches all shipments for a given warehouse ID and destination
func GetShipmentsByWarehouseIDAndDestination(warehouseID uint, destination string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("warehouse_id = ? AND destination = ?", warehouseID, destination).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByWarehouseIDAndDestinationAndStatus fetches all shipments for a given warehouse ID, destination, and status
func GetShipmentsByWarehouseIDAndDestinationAndStatus(warehouseID uint, destination, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("warehouse_id = ? AND destination = ? AND status = ?", warehouseID, destination, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByCarrierAndOrigin fetches all shipments for a given carrier and origin
func GetShipmentsByCarrierAndOrigin(carrier, origin string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("carrier = ? AND origin = ?", carrier, origin).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByCarrierAndOriginAndStatus fetches all shipments for a given carrier, origin, and status
func GetShipmentsByCarrierAndOriginAndStatus(carrier, origin, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("carrier = ? AND origin = ? AND status = ?", carrier, origin, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByCarrierAndDestination fetches all shipments for a given carrier and destination
func GetShipmentsByCarrierAndDestination(carrier, destination string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("carrier = ? AND destination = ?", carrier, destination).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByCarrierAndDestinationAndStatus fetches all shipments for a given carrier, destination, and status
func GetShipmentsByCarrierAndDestinationAndStatus(carrier, destination, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("carrier = ? AND destination = ? AND status = ?", carrier, destination, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByOriginAndDestination fetches all shipments for a given origin and destination
func GetShipmentsByOriginAndDestination(origin, destination string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("origin = ? AND destination = ?", origin, destination).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByOriginAndDestinationAndStatus fetches all shipments for a given origin, destination, and status
func GetShipmentsByOriginAndDestinationAndStatus(origin, destination, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("origin = ? AND destination = ? AND status = ?", origin, destination, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByOriginAndDestinationAndCarrier fetches all shipments for a given origin, destination, and carrier
func GetShipmentsByOriginAndDestinationAndCarrier(origin, destination, carrier string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("origin = ? AND destination = ? AND carrier = ?", origin, destination, carrier).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByOriginAndDestinationAndCarrierAndStatus fetches all shipments for a given origin, destination, carrier, and status
func GetShipmentsByOriginAndDestinationAndCarrierAndStatus(origin, destination, carrier, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("origin = ? AND destination = ? AND carrier = ? AND status = ?", origin, destination, carrier, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByDestinationAndCarrier fetches all shipments for a given destination and carrier
func GetShipmentsByDestinationAndCarrier(destination, carrier string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("destination = ? AND carrier = ?", destination, carrier).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByDestinationAndCarrierAndStatus fetches all shipments for a given destination, carrier, and status
func GetShipmentsByDestinationAndCarrierAndStatus(destination, carrier, status string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("destination = ? AND carrier = ? AND status = ?", destination, carrier, status).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByDestinationAndOrigin fetches all shipments for a given destination and origin
func GetShipmentsByDestinationAndOrigin(destination, origin string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("destination = ? AND origin = ?", destination, origin).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}

	return shipments, nil
}

// GetShipmentsByWarehouseIDAndProductID fetches shipments by warehouse ID and product ID
func GetShipmentsByWarehouseIDAndProductID(warehouseID, productID uint) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("warehouse_id = ? AND product_id = ?", warehouseID, productID).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}
	return shipments, nil
}

// GetShipmentsByProductIDAndCarrier fetches shipments by product ID and carrier
func GetShipmentsByProductIDAndCarrier(productID uint, carrier string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("product_id = ? AND carrier = ?", productID, carrier).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}
	return shipments, nil
}

// GetShipmentsByProductIDAndDestination fetches shipments by product ID and destination
func GetShipmentsByProductIDAndDestination(productID uint, destination string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("product_id = ? AND destination = ?", productID, destination).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}
	return shipments, nil
}

// GetShipmentsByProductIDAndOrigin fetches shipments by product ID and origin
func GetShipmentsByProductIDAndOrigin(productID uint, origin string) ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Where("product_id = ? AND origin = ?", productID, origin).Find(&shipments)
	if result.Error != nil {
		return nil, result.Error
	}
	return shipments, nil
}
