package controllers

import (
	"encoding/json"
	"inventory-supply-chain-system/models"
	"inventory-supply-chain-system/services"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

// CreateOrder handles the creation of a new order
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.CreateOrder(order)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// GetOrder handles retrieving a single order by ID
func GetOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	order, err := services.GetOrder(params["id"])
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// UpdateOrder handles updating an existing order
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var order models.Order

	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	orderID, err := strconv.ParseUint(params["id"], 10, 64) // Convert params["id"] to uint64
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	order.ID = uint(orderID) // Convert orderID to uint

	err = services.UpdateOrder(order)
	if err != nil {
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// DeleteOrder handles deleting an order by ID
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := services.DeleteOrder(params["id"])
	if err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CreateOrderHandler handles creating a new order
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.CreateOrder(order)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}

// GetAllOrdersHandler returns all orders
func GetAllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := services.GetAllOrders()
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrderHandler returns an order by ID
func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	order, err := services.GetOrder(id)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

// UpdateOrderHandler updates an order by ID
func UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var order models.Order
	order, err := services.GetOrder(id)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.UpdateOrder(order)
	if err != nil {
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}

// DeleteOrderHandler deletes an order by ID
func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := services.DeleteOrder(id)
	if err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetOrdersByCustomerIDHandler returns orders by customer ID
func GetOrdersByCustomerIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customerID"]

	orders, err := services.GetOrdersByCustomerID(customerID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByVendorIDHandler returns orders by vendor ID
func GetOrdersByVendorIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vendorID := params["vendorID"]

	orders, err := services.GetOrdersByVendorID(vendorID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByProductIDHandler returns orders by product ID
func GetOrdersByProductIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["productID"]

	orders, err := services.GetOrdersByProductID(productID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByShipmentIDHandler returns orders by shipment ID
func GetOrdersByShipmentIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shipmentID := params["shipmentID"]

	orders, err := services.GetOrdersByShipmentID(shipmentID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByStatusHandler returns orders by status
func GetOrdersByStatusHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	status := params["status"]

	orders, err := services.GetOrdersByStatus(status)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByDateRangeHandler returns orders by date range
func GetOrdersByDateRangeHandler(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start")
	endDate := r.URL.Query().Get("end")

	orders, err := services.GetOrdersByDateRange(startDate, endDate)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByCustomerIDAndStatusHandler returns orders by customer ID and status
func GetOrdersByCustomerIDAndStatusHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customerID"]
	status := params["status"]

	orders, err := services.GetOrdersByCustomerIDAndStatus(customerID, status)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByVendorIDAndStatusHandler returns orders by vendor ID and status
func GetOrdersByVendorIDAndStatusHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vendorID := params["vendorID"]
	status := params["status"]

	orders, err := services.GetOrdersByVendorIDAndStatus(vendorID, status)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByProductIDAndStatusHandler returns orders by product ID and status
func GetOrdersByProductIDAndStatusHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["productID"]
	status := params["status"]

	orders, err := services.GetOrdersByProductIDAndStatus(productID, status)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByVendorIDAndProductIDHandler returns orders by vendor ID and product ID
func GetOrdersByVendorIDAndProductIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vendorID := params["vendorID"]
	productID := params["productID"]

	orders, err := services.GetOrdersByVendorIDAndProductID(vendorID, productID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByVendorIDAndShipmentIDHandler returns orders by vendor ID and shipment ID
func GetOrdersByVendorIDAndShipmentIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vendorID := params["vendorID"]
	shipmentID := params["shipmentID"]

	orders, err := services.GetOrdersByVendorIDAndShipmentID(vendorID, shipmentID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByProductIDAndShipmentIDHandler returns orders by product ID and shipment ID
func GetOrdersByProductIDAndShipmentIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["productID"]
	shipmentID := params["shipmentID"]

	orders, err := services.GetOrdersByProductIDAndShipmentID(productID, shipmentID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByCustomerIDAndProductIDHandler returns orders by customer ID and product ID
func GetOrdersByCustomerIDAndProductIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customerID"]
	productID := params["productID"]

	orders, err := services.GetOrdersByCustomerIDAndProductID(customerID, productID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByCustomerIDAndShipmentIDHandler returns orders by customer ID and shipment ID
func GetOrdersByCustomerIDAndShipmentIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customerID"]
	shipmentID := params["shipmentID"]

	orders, err := services.GetOrdersByCustomerIDAndShipmentID(customerID, shipmentID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByCustomerIDAndProductIDAndShipmentIDHandler returns orders by customer ID, product ID, and shipment ID
func GetOrdersByCustomerIDAndProductIDAndShipmentIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customerID"]
	productID := params["productID"]
	shipmentID := params["shipmentID"]

	orders, err := services.GetOrdersByCustomerIDAndProductIDAndShipmentID(customerID, productID, shipmentID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByVendorIDAndProductIDAndShipmentIDHandler returns orders by vendor ID, product ID, and shipment ID
func GetOrdersByVendorIDAndProductIDAndShipmentIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vendorID := params["vendorID"]
	productID := params["productID"]
	shipmentID := params["shipmentID"]

	orders, err := services.GetOrdersByVendorIDAndProductIDAndShipmentID(vendorID, productID, shipmentID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByCustomerIDAndVendorIDAndProductIDAndShipmentIDHandler returns orders by customer ID, vendor ID, product ID, and shipment ID
func GetOrdersByCustomerIDAndVendorIDAndProductIDAndShipmentIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customerID"]
	vendorID := params["vendorID"]
	productID := params["productID"]
	shipmentID := params["shipmentID"]

	orders, err := services.GetOrdersByCustomerIDAndVendorIDAndProductIDAndShipmentID(customerID, vendorID, productID, shipmentID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByStatusAndDateRangeHandler returns orders by status and date range
func GetOrdersByStatusAndDateRangeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	status := params["status"]
	startDate := r.URL.Query().Get("start")
	endDate := r.URL.Query().Get("end")

	orders, err := services.GetOrdersByStatusAndDateRange(status, startDate, endDate)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetOrdersByCustomerIDAndStatusAndDateRangeHandler returns orders by customer ID, status, and date range
func GetOrdersByCustomerIDAndStatusAndDateRangeHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	customerID := params["customerID"]
	status := params["status"]
	startDate := r.URL.Query().Get("start")
	endDate := r.URL.Query().Get("end")

	orders, err := services.GetOrdersByCustomerIDAndStatusAndDateRange(customerID, status, startDate, endDate)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}
