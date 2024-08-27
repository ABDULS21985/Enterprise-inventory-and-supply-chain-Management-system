package main

import (
	"inventory-supply-chain-system/controllers"
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/internal/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
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

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
