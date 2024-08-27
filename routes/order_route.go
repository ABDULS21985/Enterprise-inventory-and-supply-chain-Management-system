package routes

import (
	"inventory-supply-chain-system/controllers"

	"github.com/gorilla/mux"
)

// RegisterOrderRoutes registers order-related routes with the router
func RegisterOrderRoutes(router *mux.Router) {
	// Order CRUD operations
	router.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/orders", controllers.GetAllOrdersHandler).Methods("GET")
	router.HandleFunc("/orders/{id:[0-9]+}", controllers.GetOrder).Methods("GET")
	router.HandleFunc("/orders/{id:[0-9]+}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/orders/{id:[0-9]+}", controllers.DeleteOrder).Methods("DELETE")

	// Order filters based on different attributes
	router.HandleFunc("/orders/customer/{customerID}", controllers.GetOrdersByCustomerIDHandler).Methods("GET")
	router.HandleFunc("/orders/vendor/{vendorID}", controllers.GetOrdersByVendorIDHandler).Methods("GET")
	router.HandleFunc("/orders/product/{productID}", controllers.GetOrdersByProductIDHandler).Methods("GET")
	router.HandleFunc("/orders/shipment/{shipmentID}", controllers.GetOrdersByShipmentIDHandler).Methods("GET")
	router.HandleFunc("/orders/status/{status}", controllers.GetOrdersByStatusHandler).Methods("GET")
	router.HandleFunc("/orders/date-range", controllers.GetOrdersByDateRangeHandler).Methods("GET")

	// Combined filters
	router.HandleFunc("/orders/customer/{customerID}/status/{status}", controllers.GetOrdersByCustomerIDAndStatusHandler).Methods("GET")
	router.HandleFunc("/orders/vendor/{vendorID}/status/{status}", controllers.GetOrdersByVendorIDAndStatusHandler).Methods("GET")
	router.HandleFunc("/orders/product/{productID}/status/{status}", controllers.GetOrdersByProductIDAndStatusHandler).Methods("GET")

	router.HandleFunc("/orders/vendor/{vendorID}/product/{productID}", controllers.GetOrdersByVendorIDAndProductIDHandler).Methods("GET")
	router.HandleFunc("/orders/vendor/{vendorID}/shipment/{shipmentID}", controllers.GetOrdersByVendorIDAndShipmentIDHandler).Methods("GET")
	router.HandleFunc("/orders/product/{productID}/shipment/{shipmentID}", controllers.GetOrdersByProductIDAndShipmentIDHandler).Methods("GET")

	// Complex combined filters
	router.HandleFunc("/orders/customer/{customerID}/product/{productID}/shipment/{shipmentID}", controllers.GetOrdersByCustomerIDAndProductIDAndShipmentIDHandler).Methods("GET")
	router.HandleFunc("/orders/vendor/{vendorID}/product/{productID}/shipment/{shipmentID}", controllers.GetOrdersByVendorIDAndProductIDAndShipmentIDHandler).Methods("GET")
	router.HandleFunc("/orders/customer/{customerID}/vendor/{vendorID}/product/{productID}/shipment/{shipmentID}", controllers.GetOrdersByCustomerIDAndVendorIDAndProductIDAndShipmentIDHandler).Methods("GET")

	// Filters based on status and date range
	router.HandleFunc("/orders/status/{status}/date-range", controllers.GetOrdersByStatusAndDateRangeHandler).Methods("GET")
	router.HandleFunc("/orders/customer/{customerID}/status/{status}/date-range", controllers.GetOrdersByCustomerIDAndStatusAndDateRangeHandler).Methods("GET")
}
