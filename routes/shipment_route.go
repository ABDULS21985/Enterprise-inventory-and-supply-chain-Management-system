package routes

import (
	"inventory-supply-chain-system/controllers"

	"github.com/gorilla/mux"
)

// RegisterShipmentRoutes registers shipment-related routes with the router
func RegisterShipmentRoutes(router *mux.Router) {
	// Shipment CRUD operations
	router.HandleFunc("/shipments", controllers.CreateShipment).Methods("POST")
	router.HandleFunc("/shipments", controllers.GetShipments).Methods("GET")
	router.HandleFunc("/shipments/{id:[0-9]+}", controllers.GetShipmentByID).Methods("GET")
	router.HandleFunc("/shipments/{id:[0-9]+}", controllers.UpdateShipment).Methods("PUT")
	router.HandleFunc("/shipments/{id:[0-9]+}", controllers.DeleteShipment).Methods("DELETE")

	// Shipment filters based on attributes
	router.HandleFunc("/shipments/status", controllers.GetShipmentsByStatus).Methods("GET")
	router.HandleFunc("/shipments/product", controllers.GetShipmentsByProductID).Methods("GET")
	router.HandleFunc("/shipments/destination", controllers.GetShipmentsByDestination).Methods("GET")
	router.HandleFunc("/shipments/origin", controllers.GetShipmentsByOrigin).Methods("GET")
	router.HandleFunc("/shipments/warehouse", controllers.GetShipmentsByWarehouseID).Methods("GET")
	router.HandleFunc("/shipments/carrier", controllers.GetShipmentsByCarrier).Methods("GET")
	router.HandleFunc("/shipments/tracking", controllers.GetShipmentsByTrackingNumber).Methods("GET")

	// Combined filters
	router.HandleFunc("/shipments/warehouse/status", controllers.GetShipmentsByWarehouseIDAndStatus).Methods("GET")
	router.HandleFunc("/shipments/product/status", controllers.GetShipmentsByProductIDAndStatus).Methods("GET")
	router.HandleFunc("/shipments/carrier/status", controllers.GetShipmentsByCarrierAndStatus).Methods("GET")
	router.HandleFunc("/shipments/destination/status", controllers.GetShipmentsByDestinationAndStatus).Methods("GET")
	router.HandleFunc("/shipments/origin/status", controllers.GetShipmentsByOriginAndStatus).Methods("GET")

	// Shipment combinations with product ID
	router.HandleFunc("/shipments/warehouse/product", controllers.GetShipmentsByWarehouseIDAndProductID).Methods("GET")
	router.HandleFunc("/shipments/warehouse/carrier", controllers.GetShipmentsByWarehouseIDAndCarrier).Methods("GET")
	router.HandleFunc("/shipments/warehouse/destination", controllers.GetShipmentsByWarehouseIDAndDestination).Methods("GET")
	router.HandleFunc("/shipments/warehouse/origin", controllers.GetShipmentsByWarehouseIDAndOrigin).Methods("GET")

	router.HandleFunc("/shipments/product/carrier", controllers.GetShipmentsByProductIDAndCarrier).Methods("GET")
	router.HandleFunc("/shipments/product/destination", controllers.GetShipmentsByProductIDAndDestination).Methods("GET")
	router.HandleFunc("/shipments/product/origin", controllers.GetShipmentsByProductIDAndOrigin).Methods("GET")
}
