package routes

import (
	"inventory-supply-chain-system/controllers"

	"github.com/gorilla/mux"
)

// RegisterSupplierRoutes registers supplier-related routes with the router
func RegisterSupplierRoutes(router *mux.Router) {
	router.HandleFunc("/suppliers", controllers.CreateSupplier).Methods("POST")
	router.HandleFunc("/suppliers", controllers.GetSuppliers).Methods("GET")
	router.HandleFunc("/suppliers/{id:[0-9]+}", controllers.GetSupplierByID).Methods("GET")
	router.HandleFunc("/suppliers/{id:[0-9]+}", controllers.UpdateSupplier).Methods("PUT")
	router.HandleFunc("/suppliers/{id:[0-9]+}", controllers.DeleteSupplier).Methods("DELETE")

	router.HandleFunc("/suppliers/category", controllers.GetSuppliersByCategory).Methods("GET")
	router.HandleFunc("/suppliers/product", controllers.GetSuppliersByProductID).Methods("GET")
	router.HandleFunc("/suppliers/location", controllers.GetSuppliersByLocation).Methods("GET")
	router.HandleFunc("/suppliers/rating", controllers.GetSuppliersByRating).Methods("GET")

	router.HandleFunc("/suppliers/product-location", controllers.GetSuppliersByProductIDAndLocation).Methods("GET")
	router.HandleFunc("/suppliers/product-rating", controllers.GetSuppliersByProductIDAndRating).Methods("GET")
	router.HandleFunc("/suppliers/location-rating", controllers.GetSuppliersByLocationAndRating).Methods("GET")

	router.HandleFunc("/suppliers/product-location-rating", controllers.GetSuppliersByProductIDAndLocationAndRating).Methods("GET")
	router.HandleFunc("/suppliers/category-location-rating", controllers.GetSuppliersByCategoryAndLocationAndRating).Methods("GET")
}
