package controllers

import (
	"encoding/json"
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
	"net/http"
)

// GetProfileHandler returns the profile of the authenticated user
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the user ID from the context (set by the AuthMiddleware)
	userID := r.Context().Value("userID").(uint)

	// Fetch the user information from the database
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return user information, excluding the password
	user.Password = ""
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
