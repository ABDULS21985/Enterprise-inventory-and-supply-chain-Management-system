package routes

import (
	"inventory-supply-chain-system/controllers"
	"inventory-supply-chain-system/internal/middlewares"

	"github.com/gorilla/mux"
)

// VendorRoutes defines the vendor routes
func RegisterVendorRoutes(r *mux.Router) {
	api := r.PathPrefix("/api/vendors").Subrouter()
	api.Use(middlewares.AuthMiddleware)

	api.HandleFunc("", controllers.CreateVendor).Methods("POST")
	api.HandleFunc("/{id}", controllers.GetVendor).Methods("GET")
	api.HandleFunc("/{id}", controllers.UpdateVendor).Methods("PUT")
	api.HandleFunc("/{id}", controllers.DeleteVendor).Methods("DELETE")
}
