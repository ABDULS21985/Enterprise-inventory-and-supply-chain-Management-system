package routes

import (
	"inventory-supply-chain-system/controllers"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes defines the user routes
func RegisterUserRoutes(r *mux.Router) {
	api := r.PathPrefix("/api/users").Subrouter()

	// User registration and login
	api.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	api.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	// CRUD operations
	api.HandleFunc("/{id}", controllers.GetUserController).Methods("GET")
	api.HandleFunc("/{id}", controllers.UpdateUserController).Methods("PUT")
	api.HandleFunc("/{id}", controllers.DeleteUserController).Methods("DELETE")
	api.HandleFunc("", controllers.ListUsersController).Methods("GET")

	// Additional operations
	api.HandleFunc("/{id}/password", controllers.ChangePasswordController).Methods("PUT")
	api.HandleFunc("/reset-password", controllers.ResetPasswordController).Methods("POST")
	api.HandleFunc("/verify", controllers.VerifyUserController).Methods("POST")
	api.HandleFunc("/unverify", controllers.UnverifyUserController).Methods("POST")
	api.HandleFunc("/{id}/add-role", controllers.AddRoleController).Methods("POST")
	api.HandleFunc("/{id}/remove-role", controllers.RemoveRoleController).Methods("DELETE")
	api.HandleFunc("/{id}/add-permission", controllers.AddPermissionController).Methods("POST")
	api.HandleFunc("/{id}/remove-permission", controllers.RemovePermissionController).Methods("DELETE")
	api.HandleFunc("/{id}/add-address", controllers.AddAddressController).Methods("POST")
	api.HandleFunc("/{id}/remove-address", controllers.RemoveAddressController).Methods("DELETE")

	// Get user by email using a query parameter
	api.HandleFunc("/by-email", controllers.GetUserByEmailController).Methods("GET").Queries("email", "{email}")
}
