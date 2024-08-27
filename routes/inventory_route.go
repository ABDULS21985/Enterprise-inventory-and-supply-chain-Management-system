package routes

import (
	"inventory-supply-chain-system/controllers"
	"inventory-supply-chain-system/internal/middlewares"

	"github.com/gorilla/mux"
)

// InventoryRoutes defines the inventory-related routes
func RegisterInventoryRoutes(r *mux.Router) {
	api := r.PathPrefix("/api/inventory").Subrouter()
	api.Use(middlewares.AuthMiddleware)

	api.HandleFunc("", controllers.CreateInventory).Methods("POST")
	api.HandleFunc("", controllers.GetInventoryItems).Methods("GET")
	api.HandleFunc("/{id}", controllers.GetInventory).Methods("GET")
	api.HandleFunc("/{id}", controllers.UpdateInventory).Methods("PUT")
	api.HandleFunc("/{id}", controllers.DeleteInventory).Methods("DELETE")
	api.HandleFunc("/product/{productID}", controllers.GetInventoryByProductID).Methods("GET")
	api.HandleFunc("/warehouse/{warehouseID}", controllers.GetInventoryByWarehouseID).Methods("GET")
}
