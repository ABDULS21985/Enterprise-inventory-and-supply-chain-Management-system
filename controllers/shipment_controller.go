package controllers

import (
	"encoding/json"
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
	"net/http"
)

func GetShipments(w http.ResponseWriter, r *http.Request) {
	var shipments []models.Shipment
	db.DB.Find(&shipments)
	json.NewEncoder(w).Encode(shipments)
}

// GetShipment returns a shipment by ID

func GetShipment(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var shipment models.Shipment
	db.DB.First(&shipment, id)
	json.NewEncoder(w).Encode(shipment)
}

// CreateShipment creates a new shipment
func CreateShipment(w http.ResponseWriter, r *http.Request) {
	var shipment models.Shipment
	err := json.NewDecoder(r.Body).Decode(&shipment)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&shipment)
	if result.Error != nil {
		http.Error(w, "Failed to create shipment", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipment)
}

// UpdateShipment updates an existing shipment
func UpdateShipment(w http.ResponseWriter, r *http.Request) {
	var shipment models.Shipment
	err := json.NewDecoder(r.Body).Decode(&shipment)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Save(&shipment)
	if result.Error != nil {
		http.Error(w, "Failed to update shipment", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipment)
}

// DeleteShipment deletes a shipment by ID
func DeleteShipment(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var shipment models.Shipment
	db.DB.Delete(&shipment, id)
	json.NewEncoder(w).Encode("Shipment deleted successfully")
}
