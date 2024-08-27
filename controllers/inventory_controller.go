package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"inventory-supply-chain-system/models"
	"inventory-supply-chain-system/services"

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

	err = services.CreateInventoryItem(inventory)
	if err != nil {
		http.Error(w, "Failed to create inventory", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(inventory)
}

// GetInventoryItems fetches all inventory items
func GetInventoryItems(w http.ResponseWriter, r *http.Request) {
	inventoryItems, err := services.GetInventoryItems()
	if err != nil {
		http.Error(w, "Failed to retrieve inventory items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventoryItems)
}

// GetInventory fetches an inventory item by its ID
func GetInventory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	inventoryItem, err := services.GetInventoryItemByID(uint(id))
	if err != nil {
		http.Error(w, "Inventory item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventoryItem)
}

// UpdateInventory updates an existing inventory item
func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	inventoryItem, err := services.GetInventoryItemByID(uint(id))
	if err != nil {
		http.Error(w, "Inventory item not found", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&inventoryItem)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.UpdateInventoryItem(inventoryItem)
	if err != nil {
		http.Error(w, "Failed to update inventory", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventoryItem)
}

// DeleteInventory deletes an inventory item
func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteInventoryItem(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete inventory", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetInventoryByProductID fetches all inventory items for a given product ID
func GetInventoryByProductID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["productID"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	inventoryItems, err := services.GetInventoryItemsByProductID(uint(productID))
	if err != nil {
		http.Error(w, "Failed to retrieve inventory items by product ID", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventoryItems)
}

// GetInventoryByWarehouseID fetches all inventory items for a given warehouse ID
func GetInventoryByWarehouseID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	warehouseID, err := strconv.Atoi(vars["warehouseID"])
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	inventoryItems, err := services.GetInventoryItemsByWarehouseID(uint(warehouseID))
	if err != nil {
		http.Error(w, "Failed to retrieve inventory items by warehouse ID", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventoryItems)
}
