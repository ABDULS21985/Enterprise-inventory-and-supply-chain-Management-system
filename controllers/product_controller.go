package controllers

import (
	"encoding/json"
	"net/http"

	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
)

// GetProducts returns a list of all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	db.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

// GetProduct returns a product by ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var product models.Product
	db.DB.First(&product, id)
	json.NewEncoder(w).Encode(product)
}

// CreateProduct creates a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&product)
	if result.Error != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// UpdateProduct updates an existing product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Save(&product)
	if result.Error != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// DeleteProduct deletes a product by ID
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	var product models.Product
	db.DB.Delete(&product, id)
}
