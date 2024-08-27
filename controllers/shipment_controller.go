package controllers

import (
	"encoding/json"
	"inventory-supply-chain-system/models"
	"inventory-supply-chain-system/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateShipment handles the creation of a new shipment
func CreateShipment(w http.ResponseWriter, r *http.Request) {
	var shipment models.Shipment
	err := json.NewDecoder(r.Body).Decode(&shipment)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.CreateShipment(shipment)
	if err != nil {
		http.Error(w, "Error creating shipment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shipment)
}

// GetShipments fetches all shipments
func GetShipments(w http.ResponseWriter, r *http.Request) {
	shipments, err := services.GetShipments()
	if err != nil {
		http.Error(w, "Error fetching shipments", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentByID fetches a shipment by its ID
func GetShipmentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid shipment ID", http.StatusBadRequest)
		return
	}

	shipment, err := services.GetShipmentByID(uint(id))
	if err != nil {
		http.Error(w, "Shipment not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(shipment)
}

// UpdateShipment updates an existing shipment
func UpdateShipment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid shipment ID", http.StatusBadRequest)
		return
	}

	var shipment models.Shipment
	err = json.NewDecoder(r.Body).Decode(&shipment)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	shipment.ID = uint(id)
	err = services.UpdateShipment(&shipment)
	if err != nil {
		http.Error(w, "Error updating shipment", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipment)
}

// DeleteShipment deletes a shipment by its ID
func DeleteShipment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid shipment ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteShipment(uint(id))
	if err != nil {
		http.Error(w, "Error deleting shipment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetShipmentsByStatus fetches all shipments by status
func GetShipmentsByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status == "" {
		http.Error(w, "Missing status parameter", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByStatus(status)
	if err != nil {
		http.Error(w, "Error fetching shipments by status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByProductID fetches all shipments by product ID
func GetShipmentsByProductID(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productIDStr == "" {
		http.Error(w, "Invalid or missing product ID", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByProductID(uint(productID))
	if err != nil {
		http.Error(w, "Error fetching shipments by product ID", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByDestination fetches all shipments by destination
func GetShipmentsByDestination(w http.ResponseWriter, r *http.Request) {
	destination := r.URL.Query().Get("destination")
	if destination == "" {
		http.Error(w, "Missing destination parameter", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByDestination(destination)
	if err != nil {
		http.Error(w, "Error fetching shipments by destination", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByOrigin fetches all shipments by origin
func GetShipmentsByOrigin(w http.ResponseWriter, r *http.Request) {
	origin := r.URL.Query().Get("origin")
	if origin == "" {
		http.Error(w, "Missing origin parameter", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByOrigin(origin)
	if err != nil {
		http.Error(w, "Error fetching shipments by origin", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByWarehouseID fetches all shipments by warehouse ID
func GetShipmentsByWarehouseID(w http.ResponseWriter, r *http.Request) {
	warehouseIDStr := r.URL.Query().Get("warehouse_id")
	warehouseID, err := strconv.Atoi(warehouseIDStr)
	if err != nil || warehouseIDStr == "" {
		http.Error(w, "Invalid or missing warehouse ID", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByWarehouseID(uint(warehouseID))
	if err != nil {
		http.Error(w, "Error fetching shipments by warehouse ID", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByCarrier fetches all shipments by carrier
func GetShipmentsByCarrier(w http.ResponseWriter, r *http.Request) {
	carrier := r.URL.Query().Get("carrier")
	if carrier == "" {
		http.Error(w, "Missing carrier parameter", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByCarrier(carrier)
	if err != nil {
		http.Error(w, "Error fetching shipments by carrier", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByTrackingNumber fetches a shipment by its tracking number
func GetShipmentsByTrackingNumber(w http.ResponseWriter, r *http.Request) {
	trackingNumber := r.URL.Query().Get("tracking_number")
	if trackingNumber == "" {
		http.Error(w, "Missing tracking number parameter", http.StatusBadRequest)
		return
	}

	shipment, err := services.GetShipmentsByTrackingNumber(trackingNumber)
	if err != nil {
		http.Error(w, "Error fetching shipment by tracking number", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipment)
}

// GetShipmentsByWarehouseIDAndStatus fetches all shipments by warehouse ID and status
func GetShipmentsByWarehouseIDAndStatus(w http.ResponseWriter, r *http.Request) {
	warehouseIDStr := r.URL.Query().Get("warehouse_id")
	status := r.URL.Query().Get("status")

	warehouseID, err := strconv.Atoi(warehouseIDStr)
	if err != nil || warehouseIDStr == "" || status == "" {
		http.Error(w, "Invalid or missing parameters", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByWarehouseIDAndStatus(uint(warehouseID), status)
	if err != nil {
		http.Error(w, "Error fetching shipments by warehouse ID and status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByProductIDAndStatus fetches all shipments by product ID and status
func GetShipmentsByProductIDAndStatus(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("product_id")
	status := r.URL.Query().Get("status")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productIDStr == "" || status == "" {
		http.Error(w, "Invalid or missing parameters", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByProductIDAndStatus(uint(productID), status)
	if err != nil {
		http.Error(w, "Error fetching shipments by product ID and status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByCarrierAndStatus fetches all shipments by carrier and status
func GetShipmentsByCarrierAndStatus(w http.ResponseWriter, r *http.Request) {
	carrier := r.URL.Query().Get("carrier")
	status := r.URL.Query().Get("status")

	if carrier == "" || status == "" {
		http.Error(w, "Missing carrier or status parameter", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByCarrierAndStatus(carrier, status)
	if err != nil {
		http.Error(w, "Error fetching shipments by carrier and status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByDestinationAndStatus fetches all shipments by destination and status
func GetShipmentsByDestinationAndStatus(w http.ResponseWriter, r *http.Request) {
	destination := r.URL.Query().Get("destination")
	status := r.URL.Query().Get("status")

	if destination == "" || status == "" {
		http.Error(w, "Missing destination or status parameter", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByDestinationAndStatus(destination, status)
	if err != nil {
		http.Error(w, "Error fetching shipments by destination and status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByOriginAndStatus fetches all shipments by origin and status
func GetShipmentsByOriginAndStatus(w http.ResponseWriter, r *http.Request) {
	origin := r.URL.Query().Get("origin")
	status := r.URL.Query().Get("status")

	if origin == "" || status == "" {
		http.Error(w, "Missing origin or status parameter", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByOriginAndStatus(origin, status)
	if err != nil {
		http.Error(w, "Error fetching shipments by origin and status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByWarehouseIDAndProductID fetches all shipments by warehouse ID and product ID
func GetShipmentsByWarehouseIDAndProductID(w http.ResponseWriter, r *http.Request) {
	warehouseIDStr := r.URL.Query().Get("warehouse_id")
	productIDStr := r.URL.Query().Get("product_id")

	warehouseID, _ := strconv.Atoi(warehouseIDStr)
	productID, _ := strconv.Atoi(productIDStr)
	if warehouseIDStr == "" || productIDStr == "" {
		http.Error(w, "Invalid or missing parameters", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByWarehouseIDAndProductID(uint(warehouseID), uint(productID))
	if err != nil {
		http.Error(w, "Error fetching shipments by warehouse ID and product ID", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByWarehouseIDAndCarrier fetches all shipments by warehouse ID and carrier
func GetShipmentsByWarehouseIDAndCarrier(w http.ResponseWriter, r *http.Request) {
	warehouseIDStr := r.URL.Query().Get("warehouse_id")
	carrier := r.URL.Query().Get("carrier")

	warehouseID, err := strconv.Atoi(warehouseIDStr)
	if err != nil || warehouseIDStr == "" || carrier == "" {
		http.Error(w, "Invalid or missing parameters", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByWarehouseIDAndCarrier(uint(warehouseID), carrier)
	if err != nil {
		http.Error(w, "Error fetching shipments by warehouse ID and carrier", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByWarehouseIDAndDestination fetches all shipments by warehouse ID and destination
func GetShipmentsByWarehouseIDAndDestination(w http.ResponseWriter, r *http.Request) {
	warehouseIDStr := r.URL.Query().Get("warehouse_id")
	destination := r.URL.Query().Get("destination")

	warehouseID, err := strconv.Atoi(warehouseIDStr)
	if err != nil || warehouseIDStr == "" || destination == "" {
		http.Error(w, "Invalid or missing parameters", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByWarehouseIDAndDestination(uint(warehouseID), destination)
	if err != nil {
		http.Error(w, "Error fetching shipments by warehouse ID and destination", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByWarehouseIDAndOrigin fetches all shipments by warehouse ID and origin
func GetShipmentsByWarehouseIDAndOrigin(w http.ResponseWriter, r *http.Request) {
	warehouseIDStr := r.URL.Query().Get("warehouse_id")
	origin := r.URL.Query().Get("origin")

	warehouseID, err := strconv.Atoi(warehouseIDStr)
	if err != nil || warehouseIDStr == "" || origin == "" {
		http.Error(w, "Invalid or missing parameters", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByWarehouseIDAndOrigin(uint(warehouseID), origin)
	if err != nil {
		http.Error(w, "Error fetching shipments by warehouse ID and origin", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByProductIDAndCarrier fetches all shipments by product ID and carrier
func GetShipmentsByProductIDAndCarrier(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("product_id")
	carrier := r.URL.Query().Get("carrier")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productIDStr == "" || carrier == "" {
		http.Error(w, "Invalid or missing parameters", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByProductIDAndCarrier(uint(productID), carrier)
	if err != nil {
		http.Error(w, "Error fetching shipments by product ID and carrier", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByProductIDAndDestination fetches all shipments by product ID and destination
func GetShipmentsByProductIDAndDestination(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("product_id")
	destination := r.URL.Query().Get("destination")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productIDStr == "" || destination == "" {
		http.Error(w, "Invalid or missing parameters", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByProductIDAndDestination(uint(productID), destination)
	if err != nil {
		http.Error(w, "Error fetching shipments by product ID and destination", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipmentsByProductIDAndOrigin fetches all shipments by product ID and origin
func GetShipmentsByProductIDAndOrigin(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("product_id")
	origin := r.URL.Query().Get("origin")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productIDStr == "" || origin == "" {
		http.Error(w, "Invalid or missing parameters", http.StatusBadRequest)
		return
	}

	shipments, err := services.GetShipmentsByProductIDAndOrigin(uint(productID), origin)
	if err != nil {
		http.Error(w, "Error fetching shipments by product ID and origin", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

// GetShipment handles the HTTP GET request to fetch a shipment by its ID
func GetShipment(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL parameters
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert the ID string to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Fetch the shipment from the service
	shipment, err := services.GetShipmentByID(uint(id))
	if err != nil {
		http.Error(w, "Shipment not found", http.StatusNotFound)
		return
	}

	// Return the shipment as a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shipment)
}
