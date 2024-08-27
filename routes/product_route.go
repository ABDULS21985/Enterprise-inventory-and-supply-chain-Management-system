package routes

import (
	"inventory-supply-chain-system/controllers"

	"github.com/gorilla/mux"
)

// RegisterProductRoutes registers product-related routes with the router
func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", controllers.CreateProductHandler).Methods("POST")
	router.HandleFunc("/products", controllers.GetProductsHandler).Methods("GET")
	router.HandleFunc("/products/{id:[0-9]+}", controllers.GetProductByIDHandler).Methods("GET")
	router.HandleFunc("/products/{id:[0-9]+}", controllers.UpdateProductHandler).Methods("PUT")
	router.HandleFunc("/products/{id:[0-9]+}", controllers.DeleteProductHandler).Methods("DELETE")
	router.HandleFunc("/products/category/{category}", controllers.GetProductsByCategoryHandler).Methods("GET")
	router.HandleFunc("/products/price-range", controllers.GetProductsByPriceRangeHandler).Methods("GET")
	router.HandleFunc("/products/stock/{stock:[0-9]+}", controllers.GetProductsByStockHandler).Methods("GET")
	router.HandleFunc("/products/stock-range", controllers.GetProductsByStockRangeHandler).Methods("GET")
	router.HandleFunc("/products/category/{category}/price-range", controllers.GetProductsByCategoryAndPriceRangeHandler).Methods("GET")
	router.HandleFunc("/products/category/{category}/stock/{stock:[0-9]+}", controllers.GetProductsByCategoryAndStockHandler).Methods("GET")
	router.HandleFunc("/products/category/{category}/stock-range", controllers.GetProductsByCategoryAndStockRangeHandler).Methods("GET")
	router.HandleFunc("/products/price-stock", controllers.GetProductsByPriceRangeAndStockHandler).Methods("GET")
	router.HandleFunc("/products/price-stock-range", controllers.GetProductsByPriceRangeAndStockRangeHandler).Methods("GET")
	router.HandleFunc("/products/category-price-stock", controllers.GetProductsByCategoryPriceRangeAndStockHandler).Methods("GET")
	router.HandleFunc("/products/category-price-stock-range", controllers.GetProductsByCategoryPriceRangeAndStockRangeHandler).Methods("GET")
}
