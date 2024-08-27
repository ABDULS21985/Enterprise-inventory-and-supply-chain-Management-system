package controllers

import (
	"encoding/json"
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
	"net/http"
)

// GetVendors returns a list of all vendors
func GetVendors(w http.ResponseWriter, r *http.Request) {
	var vendors []models.Vendor
	db.DB.Find(&vendors)
	json.NewEncoder(w).Encode(vendors)
}

// GetVendor returns a vendor by ID
func GetVendor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var vendor models.Vendor
	db.DB.First(&vendor, id)
	json.NewEncoder(w).Encode(vendor)
}

// CreateVendor creates a new vendor
func CreateVendor(w http.ResponseWriter, r *http.Request) {
	var vendor models.Vendor
	err := json.NewDecoder(r.Body).Decode(&vendor)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&vendor)
	if result.Error != nil {
		http.Error(w, "Failed to create vendor", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(vendor)
}

// UpdateVendor updates an existing vendor
func UpdateVendor(w http.ResponseWriter, r *http.Request) {
	var vendor models.Vendor
	err := json.NewDecoder(r.Body).Decode(&vendor)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Save(&vendor)
	if result.Error != nil {
		http.Error(w, "Failed to update vendor", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(vendor)
}

// DeleteVendor deletes a vendor by ID
func DeleteVendor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var vendor models.Vendor
	db.DB.Delete(&vendor, id)
}

//
