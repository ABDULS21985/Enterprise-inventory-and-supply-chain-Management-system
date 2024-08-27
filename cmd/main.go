package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"inventory-supply-chain-system/controllers"
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/internal/middlewares"
	"inventory-supply-chain-system/routes"
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

	// Register protected routes
	routes.RegisterItemRoutes(api)
	routes.RegisterProductRoutes(api)
	routes.RegisterProfileRoutes(api)
	routes.RegisterSupplierRoutes(api)
	routes.RegisterOrderRoutes(api)
	routes.RegisterInventoryRoutes(api)
	routes.RegisterShipmentRoutes(api)
	routes.RegisterVendorRoutes(api)

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
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w, r *http.Request) {
		http.ServeFile(w, r, "static/405.html")
	})

	// Middleware to handle 500 and 503 errors with logging
	r.Use(recoverMiddleware)

	// Set up server with graceful shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Start the server in a goroutine
	go func() {
		log.Println("Server is running on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish the request it is handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

// recoverMiddleware handles panics and recovers with appropriate error pages.
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Recovered from panic: %v", rec)
				http.ServeFile(w, r, "static/500.html")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
