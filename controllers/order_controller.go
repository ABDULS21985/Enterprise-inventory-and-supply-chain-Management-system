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
