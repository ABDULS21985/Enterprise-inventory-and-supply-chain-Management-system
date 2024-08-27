package controllers

import (
	"encoding/json"
	"inventory-supply-chain-system/models"
	"inventory-supply-chain-system/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateSupplier handles the creation of a new supplier.
func CreateSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier models.Supplier
	err := json.NewDecoder(r.Body).Decode(&supplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.CreateSupplier(supplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(supplier)
}

// GetSuppliers handles fetching all suppliers from the database.
func GetSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := services.GetSuppliers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suppliers)
}

// GetSupplierByID handles fetching a single supplier by its ID.
func GetSupplierByID(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	supplier, err := services.GetSupplierByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(supplier)
}

// UpdateSupplier handles updating an existing supplier.
func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var supplier models.Supplier
	err = json.NewDecoder(r.Body).Decode(&supplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	supplier.ID = uint(id)
	err = services.UpdateSupplier(&supplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(supplier)
}

// DeleteSupplier handles the deletion of a supplier by its ID.
func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.DeleteSupplier(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetSuppliersByCategory handles fetching suppliers by their category.
func GetSuppliersByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	suppliers, err := services.GetSuppliersByCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suppliers)
}

// GetSuppliersByProductID handles fetching suppliers based on the products they supply.
func GetSuppliersByProductID(w http.ResponseWriter, r *http.Request) {
	productIDParam := r.URL.Query().Get("product_id")
	productID, err := strconv.ParseUint(productIDParam, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	suppliers, err := services.GetSuppliersByProductID(uint(productID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suppliers)
}

// GetSuppliersByLocation handles fetching suppliers based on their location.
func GetSuppliersByLocation(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	suppliers, err := services.GetSuppliersByLocation(location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suppliers)
}

// GetSuppliersByRating handles fetching suppliers based on their rating.
func GetSuppliersByRating(w http.ResponseWriter, r *http.Request) {
	ratingParam := r.URL.Query().Get("rating")
	rating, err := strconv.ParseFloat(ratingParam, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	suppliers, err := services.GetSuppliersByRating(float32(rating))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suppliers)
}

// GetSuppliersByProductIDAndLocation handles fetching suppliers based on the products they supply and their location.
func GetSuppliersByProductIDAndLocation(w http.ResponseWriter, r *http.Request) {
	productIDParam := r.URL.Query().Get("product_id")
	productID, err := strconv.ParseUint(productIDParam, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	location := r.URL.Query().Get("location")
	suppliers, err := services.GetSuppliersByProductIDAndLocation(uint(productID), location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suppliers)
}

// GetSuppliersByProductIDAndRating handles fetching suppliers based on the products they supply and their rating.
func GetSuppliersByProductIDAndRating(w http.ResponseWriter, r *http.Request) {
	productIDParam := r.URL.Query().Get("product_id")
	productID, err := strconv.ParseUint(productIDParam, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ratingParam := r.URL.Query().Get("rating")
	rating, err := strconv.ParseFloat(ratingParam, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	suppliers, err := services.GetSuppliersByProductIDAndRating(uint(productID), float32(rating))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suppliers)
}

// GetSuppliersByLocationAndRating handles fetching suppliers based on their location and rating.
func GetSuppliersByLocationAndRating(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	ratingParam := r.URL.Query().Get("rating")
	rating, err := strconv.ParseFloat(ratingParam, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	suppliers, err := services.GetSuppliersByLocationAndRating(location, float32(rating))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suppliers)
}

// GetSuppliersByProductIDAndLocationAndRating handles fetching suppliers based on the products they supply, their location, and rating.
func GetSuppliersByProductIDAndLocationAndRating(w http.ResponseWriter, r *http.Request) {
	productIDParam := r.URL.Query().Get("product_id")
	productID, err := strconv.ParseUint(productIDParam, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	location := r.URL.Query().Get("location")
	ratingParam := r.URL.Query().Get("rating")
	rating, err := strconv.ParseFloat(ratingParam, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	suppliers, err := services.GetSuppliersByProductIDAndLocationAndRating(uint(productID), location, float32(rating))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suppliers)
}

// GetSuppliersByCategoryAndLocationAndRating handles fetching suppliers based on their category, location, and rating.
func GetSuppliersByCategoryAndLocationAndRating(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	location := r.URL.Query().Get("location")
	ratingParam := r.URL.Query().Get("rating")
	rating, err := strconv.ParseFloat(ratingParam, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	suppliers, err := services.GetSuppliersByCategoryAndLocationAndRating(category, location, float32(rating))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suppliers)
}
