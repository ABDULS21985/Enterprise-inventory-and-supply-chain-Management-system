package controllers

import (
	"encoding/json"
	"net/http"

	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
)

func GetInventoryItems(w http.ResponseWriter, r *http.Request) {
	var inventoryItems []models.Inventory
	db.DB.Find(&inventoryItems)

	json.NewEncoder(w).Encode(inventoryItems)
}

func CreateInventoryItem(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&inventory)
	if result.Error != nil {
		http.Error(w, "Error saving inventory", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(inventory)
}
