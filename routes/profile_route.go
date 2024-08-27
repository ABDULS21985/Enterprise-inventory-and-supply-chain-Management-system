package routes

import (
	"inventory-supply-chain-system/controllers"

	"github.com/gorilla/mux"
)

// RegisterProfileRoutes registers routes for profile-related operations
func RegisterProfileRoutes(router *mux.Router) {
	router.HandleFunc("/profile", controllers.GetProfileHandler).Methods("GET")
	router.HandleFunc("/profile", controllers.UpdateProfileHandler).Methods("PUT")
	router.HandleFunc("/profile/password", controllers.UpdatePasswordHandler).Methods("PUT")
}
