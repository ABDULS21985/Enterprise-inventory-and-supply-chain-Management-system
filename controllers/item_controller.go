package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"inventory-supply-chain-system/models"
	"inventory-supply-chain-system/services"

	"github.com/gorilla/mux"
)

// CreateItem creates a new item in the inventory.
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := services.CreateItem(item); err != nil {
		http.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// GetItems fetches all items from the inventory.
func GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := services.GetItems()
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// GetItemByID fetches a single item by its ID.
func GetItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	item, err := services.GetItemByID(uint(id))
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

// UpdateItem updates an existing item in the inventory.
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := services.UpdateItem(&item); err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

// DeleteItem removes an item from the inventory by its ID.
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteItem(uint(id)); err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetItemsByCategory fetches items based on their category.
func GetItemsByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	items, err := services.GetItemsByCategory(category)
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// GetItemsByWarehouseID fetches items stored in a specific warehouse.
func GetItemsByWarehouseID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	warehouseID, err := strconv.ParseUint(vars["warehouseID"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	items, err := services.GetItemsByWarehouseID(uint(warehouseID))
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// GetItemsBySupplierID fetches items based on their supplier ID.
func GetItemsBySupplierID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	supplierID, err := strconv.ParseUint(vars["supplierID"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid supplier ID", http.StatusBadRequest)
		return
	}

	items, err := services.GetItemsBySupplierID(uint(supplierID))
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// GetItemsByStockRange fetches items within a specific stock level range.
func GetItemsByStockRange(w http.ResponseWriter, r *http.Request) {
	minStock, err := strconv.Atoi(r.URL.Query().Get("minStock"))
	if err != nil {
		http.Error(w, "Invalid minimum stock", http.StatusBadRequest)
		return
	}

	maxStock, err := strconv.Atoi(r.URL.Query().Get("maxStock"))
	if err != nil {
		http.Error(w, "Invalid maximum stock", http.StatusBadRequest)
		return
	}

	items, err := services.GetItemsByStockRange(minStock, maxStock)
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// GetItemsByPriceRange fetches items within a specific price range.
func GetItemsByPriceRange(w http.ResponseWriter, r *http.Request) {
	minPrice, err := strconv.ParseFloat(r.URL.Query().Get("minPrice"), 64)
	if err != nil {
		http.Error(w, "Invalid minimum price", http.StatusBadRequest)
		return
	}

	maxPrice, err := strconv.ParseFloat(r.URL.Query().Get("maxPrice"), 64)
	if err != nil {
		http.Error(w, "Invalid maximum price", http.StatusBadRequest)
		return
	}

	items, err := services.GetItemsByPriceRange(minPrice, maxPrice)
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// GetItemsByCategoryAndPriceRange fetches items based on their category and price range.
func GetItemsByCategoryAndPriceRange(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	minPrice, err := strconv.ParseFloat(r.URL.Query().Get("minPrice"), 64)
	if err != nil {
		http.Error(w, "Invalid minimum price", http.StatusBadRequest)
		return
	}

	maxPrice, err := strconv.ParseFloat(r.URL.Query().Get("maxPrice"), 64)
	if err != nil {
		http.Error(w, "Invalid maximum price", http.StatusBadRequest)
		return
	}

	items, err := services.GetItemsByCategoryAndPriceRange(category, minPrice, maxPrice)
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// GetItemsByCategoryAndStockRange fetches items based on their category and stock level range.
func GetItemsByCategoryAndStockRange(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	minStock, err := strconv.Atoi(r.URL.Query().Get("minStock"))
	if err != nil {
		http.Error(w, "Invalid minimum stock", http.StatusBadRequest)
		return
	}

	maxStock, err := strconv.Atoi(r.URL.Query().Get("maxStock"))
	if err != nil {
		http.Error(w, "Invalid maximum stock", http.StatusBadRequest)
		return
	}

	items, err := services.GetItemsByCategoryAndStockRange(category, minStock, maxStock)
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// GetItemsByCategoryAndSupplierID fetches items based on their category and supplier ID.
func GetItemsByCategoryAndSupplierID(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	supplierID, err := strconv.ParseUint(r.URL.Query().Get("supplierID"), 10, 32)
	if err != nil {
		http.Error(w, "Invalid supplier ID", http.StatusBadRequest)
		return
	}

	items, err := services.GetItemsByCategoryAndSupplierID(category, uint(supplierID))
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// GetItemsByCategoryAndWarehouseID fetches items based on their category and warehouse ID.
func GetItemsByCategoryAndWarehouseID(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	warehouseID, err := strconv.ParseUint(r.URL.Query().Get("warehouseID"), 10, 32)
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	items, err := services.GetItemsByCategoryAndWarehouseID(category, uint(warehouseID))
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// GetItemsBySupplierIDAndWarehouseID fetches items based on their supplier ID and warehouse ID.
func GetItemsBySupplierIDAndWarehouseID(w http.ResponseWriter, r *http.Request) {
	supplierID, err := strconv.ParseUint(r.URL.Query().Get("supplierID"), 10, 32)
	if err != nil {
		http.Error(w, "Invalid supplier ID", http.StatusBadRequest)
		return
	}

	warehouseID, err := strconv.ParseUint(r.URL.Query().Get("warehouseID"), 10, 32)
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	items, err := services.GetItemsBySupplierIDAndWarehouseID(uint(supplierID), uint(warehouseID))
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}
