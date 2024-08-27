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

// GetAllOrders returns a list of all orders
func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Find(&orders).Error
	return orders, err
}

// GetOrder returns an order by ID
func GetOrder(id string) (models.Order, error) {
	var order models.Order
	err := db.DB.First(&order, id).Error
	return order, err
}

// UpdateOrder updates an existing order
func UpdateOrder(order models.Order) error {
	result := db.DB.Save(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteOrder deletes an order by ID
func DeleteOrder(id string) error {

	result := db.DB.Where("id = ?", id).Delete(&models.Order{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetOrdersByCustomerID returns a list of orders by customer ID

func GetOrdersByCustomerID(customerID string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("customer_id = ?", customerID).Find(&orders).Error
	return orders, err
}

// GetOrdersByVendorID returns a list of orders by vendor ID
func GetOrdersByVendorID(vendorID string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("vendor_id = ?", vendorID).Find(&orders).Error
	return orders, err
}

// GetOrdersByProductID returns a list of orders by product ID
func GetOrdersByProductID(productID string) ([]models.Order, error) {

	var orders []models.Order
	err := db.DB.Where("product_id = ?", productID).Find(&orders).Error
	return orders, err
}

// GetOrdersByShipmentID returns a list of orders by shipment ID
func GetOrdersByShipmentID(shipmentID string) ([]models.Order, error) {

	var orders []models.Order
	err := db.DB.Where("shipment_id = ?", shipmentID).Find(&orders).Error
	return orders, err
}

// GetOrdersByStatus returns a list of orders by status
func GetOrdersByStatus(status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("status = ?", status).Find(&orders).Error
	return orders, err
}

// GetOrdersByDateRange returns a list of orders by date range
func GetOrdersByDateRange(startDate string, endDate string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&orders).Error
	return orders, err
}

// GetOrdersByCustomerIDAndStatus returns a list of orders by customer ID and status
func GetOrdersByCustomerIDAndStatus(customerID string, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("customer_id = ? AND status = ?", customerID, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByVendorIDAndStatus returns a list of orders by vendor ID and status
func GetOrdersByVendorIDAndStatus(vendorID string, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("vendor_id = ? AND status = ?", vendorID, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByProductIDAndStatus returns a list of orders by product ID and status
func GetOrdersByProductIDAndStatus(productID string, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("product_id = ? AND status = ?", productID, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByShipmentIDAndStatus returns a list of orders by shipment ID and status
func GetOrdersByShipmentIDAndStatus(shipmentID string, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("shipment_id = ? AND status = ?", shipmentID, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByDateRangeAndStatus returns a list of orders by date range and status
func GetOrdersByDateRangeAndStatus(startDate string, endDate string, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("created_at BETWEEN ? AND ? AND status = ?", startDate, endDate, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByCustomerIDAndDateRange returns a list of orders by customer ID and date range
func GetOrdersByCustomerIDAndDateRange(customerID string, startDate string, endDate string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("customer_id = ? AND created_at BETWEEN ? AND ?", customerID, startDate, endDate).Find(&orders).Error
	return orders, err
}

// GetOrdersByVendorIDAndDateRange returns a list of orders by vendor ID and date range
func GetOrdersByVendorIDAndDateRange(vendorID string, startDate string, endDate string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("vendor_id = ? AND created_at BETWEEN ? AND ?", vendorID, startDate, endDate).Find(&orders).Error
	return orders, err
}

// GetOrdersByProductIDAndDateRange returns a list of orders by product ID and date range
func GetOrdersByProductIDAndDateRange(productID string, startDate string, endDate string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("product_id = ? AND created_at BETWEEN ? AND ?", productID, startDate, endDate).Find(&orders).Error
	return orders, err
}

// GetOrdersByShipmentIDAndDateRange returns a list of orders by shipment ID and date range
func GetOrdersByShipmentIDAndDateRange(shipmentID string, startDate string, endDate string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("shipment_id = ? AND created_at BETWEEN ? AND ?", shipmentID, startDate, endDate).Find(&orders).Error
	return orders, err
}

// GetOrdersByCustomerIDVendorIDAndStatus returns a list of orders by customer ID, vendor ID, and status
func GetOrdersByCustomerIDVendorIDAndStatus(customerID string, vendorID string, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("customer_id = ? AND vendor_id = ? AND status = ?", customerID, vendorID, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByCustomerIDProductIDAndStatus returns a list of orders by customer ID, product ID, and status
func GetOrdersByCustomerIDProductIDAndStatus(customerID string, productID string, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("customer_id = ? AND product_id = ? AND status = ?", customerID, productID, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByCustomerIDShipmentIDAndStatus returns a list of orders by customer ID, shipment ID, and status
func GetOrdersByCustomerIDShipmentIDAndStatus(customerID string, shipmentID string, status string) ([]models.Order, error) {

	var orders []models.Order
	err := db.DB.Where("customer_id = ? AND shipment_id = ? AND status = ?", customerID, shipmentID, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByCustomerIDDateRangeAndStatus returns a list of orders by customer ID, date range, and status
func GetOrdersByCustomerIDDateRangeAndStatus(customerID string, startDate string, endDate string, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("customer_id = ? AND created_at BETWEEN ? AND ? AND status = ?", customerID, startDate, endDate, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByVendorIDProductIDAndStatus returns a list of orders by vendor ID, product ID, and status
func GetOrdersByVendorIDProductIDAndStatus(vendorID string, productID string, status string) ([]models.Order, error) {

	var orders []models.Order
	err := db.DB.Where("vendor_id = ? AND product_id = ? AND status = ?", vendorID, productID, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByVendorIDShipmentIDAndStatus returns a list of orders by vendor ID, shipment ID, and status
func GetOrdersByVendorIDShipmentIDAndStatus(vendorID string, shipmentID string, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("vendor_id = ? AND shipment_id = ? AND status = ?", vendorID, shipmentID, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByVendorIDDateRangeAndStatus returns a list of orders by vendor ID, date range, and status
func GetOrdersByVendorIDDateRangeAndStatus(vendorID string, startDate string, endDate string, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("vendor_id = ? AND created_at BETWEEN ? AND ? AND status = ?", vendorID, startDate, endDate, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByProductIDShipmentIDAndStatus returns a list of orders by product ID, shipment ID, and status
func GetOrdersByProductIDShipmentIDAndStatus(productID string, shipmentID string, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("product_id = ? AND shipment_id = ? AND status = ?", productID, shipmentID, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByProductIDDateRangeAndStatus returns a list of orders by product ID, date range, and status
func GetOrdersByProductIDDateRangeAndStatus(productID string, startDate string, endDate string, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("product_id = ? AND created_at BETWEEN ? AND ? AND status = ?", productID, startDate, endDate, status).Find(&orders).Error
	return orders, err
}

// GetOrdersByVendorIDAndProductID returns orders by vendor ID and product ID
func GetOrdersByVendorIDAndProductID(vendorID string, productID string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("vendor_id = ? AND product_id = ?", vendorID, productID).Find(&orders).Error
	return orders, err
}

// GetOrdersByVendorIDAndShipmentID returns orders by vendor ID and shipment ID
func GetOrdersByVendorIDAndShipmentID(vendorID string, shipmentID string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("vendor_id = ? AND shipment_id = ?", vendorID, shipmentID).Find(&orders).Error
	return orders, err
}

// GetOrdersByProductIDAndShipmentID returns orders by product ID and shipment ID
func GetOrdersByProductIDAndShipmentID(productID string, shipmentID string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("product_id = ? AND shipment_id = ?", productID, shipmentID).Find(&orders).Error
	return orders, err
}

// GetOrdersByCustomerIDAndProductID returns orders by customer ID and product ID
func GetOrdersByCustomerIDAndProductID(customerID string, productID string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("customer_id = ? AND product_id = ?", customerID, productID).Find(&orders).Error
	return orders, err
}

// GetOrdersByCustomerIDAndShipmentID returns orders by customer ID and shipment ID
func GetOrdersByCustomerIDAndShipmentID(customerID string, shipmentID string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("customer_id = ? AND shipment_id = ?", customerID, shipmentID).Find(&orders).Error
	return orders, err
}

// GetOrdersByCustomerIDAndProductIDAndShipmentID returns orders by customer ID, product ID, and shipment ID
func GetOrdersByCustomerIDAndProductIDAndShipmentID(customerID string, productID string, shipmentID string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("customer_id = ? AND product_id = ? AND shipment_id = ?", customerID, productID, shipmentID).Find(&orders).Error
	return orders, err
}

// GetOrdersByVendorIDAndProductIDAndShipmentID returns orders by vendor ID, product ID, and shipment ID
func GetOrdersByVendorIDAndProductIDAndShipmentID(vendorID string, productID string, shipmentID string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("vendor_id = ? AND product_id = ? AND shipment_id = ?", vendorID, productID, shipmentID).Find(&orders).Error
	return orders, err
}

// GetOrdersByCustomerIDAndVendorIDAndProductIDAndShipmentID returns orders by customer ID, vendor ID, product ID, and shipment ID
func GetOrdersByCustomerIDAndVendorIDAndProductIDAndShipmentID(customerID string, vendorID string, productID string, shipmentID string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("customer_id = ? AND vendor_id = ? AND product_id = ? AND shipment_id = ?", customerID, vendorID, productID, shipmentID).Find(&orders).Error
	return orders, err
}

// GetOrdersByStatusAndDateRange returns orders by status and date range
func GetOrdersByStatusAndDateRange(status string, startDate string, endDate string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("status = ? AND created_at BETWEEN ? AND ?", status, startDate, endDate).Find(&orders).Error
	return orders, err
}

// GetOrdersByCustomerIDAndStatusAndDateRange returns orders by customer ID, status, and date range
func GetOrdersByCustomerIDAndStatusAndDateRange(customerID string, status string, startDate string, endDate string) ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Where("customer_id = ? AND status = ? AND created_at BETWEEN ? AND ?", customerID, status, startDate, endDate).Find(&orders).Error
	return orders, err
}
