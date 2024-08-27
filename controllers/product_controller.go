package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"inventory-supply-chain-system/models"
	"inventory-supply-chain-system/services"

	"github.com/gorilla/mux"
)

// CreateProductHandler handles the creation of a new product
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := services.CreateProduct(product); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProductsHandler fetches all products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := services.GetProducts()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProductByIDHandler fetches a product by ID
func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := services.GetProductByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// UpdateProductHandler handles the update of an existing product
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := services.GetProductByID(uint(id))
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := services.UpdateProduct(product); err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// DeleteProductHandler deletes a product by ID
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteProduct(uint(id)); err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}

// GetProductsByCategoryHandler fetches products by category
func GetProductsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	products, err := services.GetProductsByCategory(category)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProductsByPriceRangeHandler fetches products by price range
func GetProductsByPriceRangeHandler(w http.ResponseWriter, r *http.Request) {
	minPrice, _ := strconv.ParseFloat(r.URL.Query().Get("minPrice"), 64)
	maxPrice, _ := strconv.ParseFloat(r.URL.Query().Get("maxPrice"), 64)

	products, err := services.GetProductsByPriceRange(minPrice, maxPrice)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProductsByStockHandler fetches products by stock level
func GetProductsByStockHandler(w http.ResponseWriter, r *http.Request) {
	stock, _ := strconv.Atoi(mux.Vars(r)["stock"])

	products, err := services.GetProductsByStock(stock)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProductsByStockRangeHandler fetches products by stock range
func GetProductsByStockRangeHandler(w http.ResponseWriter, r *http.Request) {
	minStock, _ := strconv.Atoi(r.URL.Query().Get("minStock"))
	maxStock, _ := strconv.Atoi(r.URL.Query().Get("maxStock"))

	products, err := services.GetProductsByStockRange(minStock, maxStock)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// CreateInventoryItemHandler handles the creation of a new inventory item
func CreateInventoryItemHandler(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	if err := json.NewDecoder(r.Body).Decode(&inventory); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := services.CreateInventoryItem(inventory); err != nil {
		http.Error(w, "Failed to create inventory", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(inventory)
}

// GetInventoryItemsHandler fetches all inventory items
func GetInventoryItemsHandler(w http.ResponseWriter, r *http.Request) {
	inventoryItems, err := services.GetInventoryItems()
	if err != nil {
		http.Error(w, "Failed to retrieve inventory items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventoryItems)
}

// GetInventoryItemHandler fetches an inventory item by its ID
func GetInventoryItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	inventoryItem, err := services.GetInventoryItemByID(uint(id))
	if err != nil {
		http.Error(w, "Inventory item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventoryItem)
}

// UpdateInventoryItemHandler updates an existing inventory item
func UpdateInventoryItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	inventoryItem, err := services.GetInventoryItemByID(uint(id))
	if err != nil {
		http.Error(w, "Inventory item not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&inventoryItem); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := services.UpdateInventoryItem(inventoryItem); err != nil {
		http.Error(w, "Failed to update inventory item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventoryItem)
}

// GetProductsByCategoryAndPriceRangeHandler fetches products by category and price range
func GetProductsByCategoryAndPriceRangeHandler(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	minPrice, _ := strconv.ParseFloat(r.URL.Query().Get("minPrice"), 64)
	maxPrice, _ := strconv.ParseFloat(r.URL.Query().Get("maxPrice"), 64)

	products, err := services.GetProductsByCategoryAndPriceRange(category, minPrice, maxPrice)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProductsByCategoryAndStockHandler fetches products by category and stock level
func GetProductsByCategoryAndStockHandler(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	stock, _ := strconv.Atoi(mux.Vars(r)["stock"])

	products, err := services.GetProductsByCategoryAndStock(category, stock)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProductsByCategoryAndStockRangeHandler fetches products by category and stock range
func GetProductsByCategoryAndStockRangeHandler(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	minStock, _ := strconv.Atoi(r.URL.Query().Get("minStock"))
	maxStock, _ := strconv.Atoi(r.URL.Query().Get("maxStock"))

	products, err := services.GetProductsByCategoryAndStockRange(category, minStock, maxStock)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProductsByPriceRangeAndStockHandler fetches products by price range and stock level
func GetProductsByPriceRangeAndStockHandler(w http.ResponseWriter, r *http.Request) {
	minPrice, _ := strconv.ParseFloat(r.URL.Query().Get("minPrice"), 64)
	maxPrice, _ := strconv.ParseFloat(r.URL.Query().Get("maxPrice"), 64)
	stock, _ := strconv.Atoi(mux.Vars(r)["stock"])

	products, err := services.GetProductsByPriceRangeAndStock(minPrice, maxPrice, stock)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProductsByPriceRangeAndStockRangeHandler fetches products by price range and stock range
func GetProductsByPriceRangeAndStockRangeHandler(w http.ResponseWriter, r *http.Request) {
	minPrice, _ := strconv.ParseFloat(r.URL.Query().Get("minPrice"), 64)
	maxPrice, _ := strconv.ParseFloat(r.URL.Query().Get("maxPrice"), 64)
	minStock, _ := strconv.Atoi(r.URL.Query().Get("minStock"))
	maxStock, _ := strconv.Atoi(r.URL.Query().Get("maxStock"))

	products, err := services.GetProductsByPriceRangeAndStockRange(minPrice, maxPrice, minStock, maxStock)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProductsByCategoryPriceRangeAndStockHandler fetches products by category, price range, and stock level
func GetProductsByCategoryPriceRangeAndStockHandler(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	minPrice, _ := strconv.ParseFloat(r.URL.Query().Get("minPrice"), 64)
	maxPrice, _ := strconv.ParseFloat(r.URL.Query().Get("maxPrice"), 64)
	stock, _ := strconv.Atoi(mux.Vars(r)["stock"])

	products, err := services.GetProductsByCategoryPriceRangeAndStock(category, minPrice, maxPrice, stock)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProductsByCategoryPriceRangeAndStockRangeHandler fetches products by category, price range, and stock range
func GetProductsByCategoryPriceRangeAndStockRangeHandler(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	minPrice, _ := strconv.ParseFloat(r.URL.Query().Get("minPrice"), 64)
	maxPrice, _ := strconv.ParseFloat(r.URL.Query().Get("maxPrice"), 64)
	minStock, _ := strconv.Atoi(r.URL.Query().Get("minStock"))
	maxStock, _ := strconv.Atoi(r.URL.Query().Get("maxStock"))

	products, err := services.GetProductsByCategoryPriceRangeAndStockRange(category, minPrice, maxPrice, minStock, maxStock)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

//Get
