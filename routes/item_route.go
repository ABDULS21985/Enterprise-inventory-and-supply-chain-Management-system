package routes

import (
	"inventory-supply-chain-system/controllers"

	"github.com/gorilla/mux"
)

// RegisterItemRoutes registers item-related routes with the router
func RegisterItemRoutes(router *mux.Router) {
	router.HandleFunc("/items", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/items", controllers.GetItems).Methods("GET")
	router.HandleFunc("/items/{id:[0-9]+}", controllers.GetItemByID).Methods("GET")
	router.HandleFunc("/items/{id:[0-9]+}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id:[0-9]+}", controllers.DeleteItem).Methods("DELETE")

	router.HandleFunc("/items/category", controllers.GetItemsByCategory).Methods("GET")
	router.HandleFunc("/items/warehouse/{warehouseID:[0-9]+}", controllers.GetItemsByWarehouseID).Methods("GET")
	router.HandleFunc("/items/supplier/{supplierID:[0-9]+}", controllers.GetItemsBySupplierID).Methods("GET")
	router.HandleFunc("/items/stock-range", controllers.GetItemsByStockRange).Methods("GET")
	router.HandleFunc("/items/price-range", controllers.GetItemsByPriceRange).Methods("GET")
	router.HandleFunc("/items/category-price", controllers.GetItemsByCategoryAndPriceRange).Methods("GET")
	router.HandleFunc("/items/category-stock", controllers.GetItemsByCategoryAndStockRange).Methods("GET")
	router.HandleFunc("/items/category-supplier", controllers.GetItemsByCategoryAndSupplierID).Methods("GET")
	router.HandleFunc("/items/category-warehouse", controllers.GetItemsByCategoryAndWarehouseID).Methods("GET")
	router.HandleFunc("/items/supplier-warehouse", controllers.GetItemsBySupplierIDAndWarehouseID).Methods("GET")
}
