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
// @Summary Create a new inventory item
// @Description Creates a new inventory item in the system
// @Tags inventory
// @Accept json
// @Produce json
// @Param inventory body models.Inventory true "Inventory data"
// @Success 201 {object} models.Inventory
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Failed to create inventory"
// @Router /inventory [post]
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
// @Summary Get all inventory items
// @Description Retrieves a list of all inventory items
// @Tags inventory
// @Produce json
// @Success 200 {array} models.Inventory
// @Failure 500 {string} string "Failed to retrieve inventory items"
// @Router /inventory [get]
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
// @Summary Get an inventory item by ID
// @Description Retrieves a specific inventory item by its ID
// @Tags inventory
// @Produce json
// @Param id path int true "Inventory ID"
// @Success 200 {object} models.Inventory
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Inventory item not found"
// @Router /inventory/{id} [get]
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
// @Summary Update an existing inventory item
// @Description Updates the details of an existing inventory item
// @Tags inventory
// @Accept json
// @Produce json
// @Param id path int true "Inventory ID"
// @Param inventory body models.Inventory true "Updated Inventory data"
// @Success 200 {object} models.Inventory
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Inventory item not found"
// @Failure 500 {string} string "Failed to update inventory"
// @Router /inventory/{id} [put]
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
// @Summary Delete an inventory item
// @Description Deletes an inventory item by its ID
// @Tags inventory
// @Param id path int true "Inventory ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Failed to delete inventory"
// @Router /inventory/{id} [delete]
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
// @Summary Get inventory items by product ID
// @Description Retrieves a list of inventory items by their associated product ID
// @Tags inventory
// @Produce json
// @Param productID path int true "Product ID"
// @Success 200 {array} models.Inventory
// @Failure 400 {string} string "Invalid product ID"
// @Failure 500 {string} string "Failed to retrieve inventory items by product ID"
// @Router /inventory/product/{productID} [get]
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
// @Summary Get inventory items by warehouse ID
// @Description Retrieves a list of inventory items by their associated warehouse ID
// @Tags inventory
// @Produce json
// @Param warehouseID path int true "Warehouse ID"
// @Success 200 {array} models.Inventory
// @Failure 400 {string} string "Invalid warehouse ID"
// @Failure 500 {string} string "Failed to retrieve inventory items by warehouse ID"
// @Router /inventory/warehouse/{warehouseID} [get]
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
