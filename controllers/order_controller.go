package controllers

import (
	"encoding/json"
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	db.DB.Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var order models.Order
	db.DB.First(&order, id)
	json.NewEncoder(w).Encode(order)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&order)
	if result.Error != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Save(&order)
	if result.Error != nil {
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var order models.Order
	db.DB.Delete(&order, id)
	json.NewEncoder(w).Encode("Order deleted successfully")
}
