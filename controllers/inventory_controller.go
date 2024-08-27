package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"

	"github.com/gorilla/mux"
)

// CreateInventory creates a new inventory item
func CreateInventory(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&inventory)
	if result.Error != nil {
		http.Error(w, "Failed to create inventory", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(inventory)
}

// GetInventory retrieves an inventory item by ID
func GetInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid inventory ID", http.StatusBadRequest)
		return
	}

	var inventory models.Inventory
	result := db.DB.First(&inventory, id)
	if result.Error != nil {
		http.Error(w, "Inventory not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(inventory)
}

// UpdateInventory updates an existing inventory item
func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid inventory ID", http.StatusBadRequest)
		return
	}

	var inventory models.Inventory
	result := db.DB.First(&inventory, id)
	if result.Error != nil {
		http.Error(w, "Inventory not found", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db.DB.Save(&inventory)
	json.NewEncoder(w).Encode(inventory)
}

// DeleteInventory deletes an inventory item
func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid inventory ID", http.StatusBadRequest)
		return
	}

	result := db.DB.Delete(&models.Inventory{}, id)
	if result.Error != nil {
		http.Error(w, "Failed to delete inventory", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
