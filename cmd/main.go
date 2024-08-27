package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"inventory-supply-chain-system/controllers"
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/internal/middlewares"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Log environment variables for debugging
	log.Println("DB_HOST:", os.Getenv("DB_HOST"))
	log.Println("DB_USER:", os.Getenv("DB_USER"))
	log.Println("DB_NAME:", os.Getenv("DB_NAME"))
	log.Println("DB_PORT:", os.Getenv("DB_PORT"))

	// Initialize the database connection
	db.ConnectDB()

	// Create a new router
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/api/auth/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/api/auth/login", controllers.LoginUser).Methods("POST")

	// Protected routes (Require authentication)
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware)
	api.HandleFunc("/api/users/me", controllers.GetProfileHandler).Methods("GET")

	// Inventory routes
	api.HandleFunc("/api/inventory", controllers.CreateInventory).Methods("POST")
	api.HandleFunc("/api/inventory/{id}", controllers.GetInventory).Methods("GET")
	api.HandleFunc("/api/inventory/{id}", controllers.UpdateInventory).Methods("PUT")
	api.HandleFunc("/api/inventory/{id}", controllers.DeleteInventory).Methods("DELETE")

	// Order routes
	api.HandleFunc("/api/orders", controllers.CreateOrder).Methods("POST")
	api.HandleFunc("/api/orders/{id}", controllers.GetOrder).Methods("GET")
	api.HandleFunc("/api/orders/{id}", controllers.UpdateOrder).Methods("PUT")
	api.HandleFunc("/api/orders/{id}", controllers.DeleteOrder).Methods("DELETE")

	// Shipment routes
	api.HandleFunc("/api/shipments", controllers.CreateShipment).Methods("POST")
	api.HandleFunc("/api/shipments/{id}", controllers.GetShipment).Methods("GET")
	api.HandleFunc("/api/shipments/{id}", controllers.UpdateShipment).Methods("PUT")
	api.HandleFunc("/api/shipments/{id}", controllers.DeleteShipment).Methods("DELETE")

	// Vendor routes
	api.HandleFunc("/api/vendors", controllers.CreateVendor).Methods("POST")
	api.HandleFunc("/api/vendors/{id}", controllers.GetVendor).Methods("GET")
	api.HandleFunc("/api/vendors/{id}", controllers.UpdateVendor).Methods("PUT")
	api.HandleFunc("/api/vendors/{id}", controllers.DeleteVendor).Methods("DELETE")

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// Handle 404
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/404.html")
	})

	// Handle 405
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/405.html")
	})

	// Middleware to handle 500 errors
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					http.ServeFile(w, r, "static/500.html")
				}
			}()
			next.ServeHTTP(w, r)
		})
	})

	// Middleware to handle 503 errors
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					http.ServeFile(w, r, "static/503.html")
				}
			}()
			next.ServeHTTP(w, r)
		})
	})

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
