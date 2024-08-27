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

// UpdateProfileHandler updates the profile of the authenticated user
func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the user ID from the context (set by the AuthMiddleware)
	userID := r.Context().Value("userID").(uint)

	// Fetch the user information from the database
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the updated user information from the request body
	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Update the user information
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Phone = updatedUser.Phone

	// Save the updated user information
	if err := db.DB.Save(&user).Error; err != nil {
		http.Error(w, "Failed to update user profile", http.StatusInternalServerError)
		return
	}

	// Return the updated user information, excluding the password
	user.Password = ""
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// UpdatePasswordHandler updates the password of the authenticated user
func UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the user ID from the context (set by the AuthMiddleware)
	userID := r.Context().Value("userID").(uint)

	// Fetch the user information from the database
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the updated password from the request body
	var passwordUpdate struct {
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&passwordUpdate); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Update the user's password
	user.Password = passwordUpdate.Password

	// Save the updated user information
	if err := db.DB.Save(&user).Error; err != nil {
		http.Error(w, "Failed to update password", http.StatusInternalServerError)
		return
	}

	// Return a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Password updated successfully")
}

// GetAllProfilesHandler returns all the profiles of the users
func GetAllProfilesHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all the user information from the database
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		http.Error(w, "Failed to fetch user profiles", http.StatusInternalServerError)
		return
	}

	// Return user information, excluding the password
	for i := range users {
		users[i].Password = ""
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// DeleteProfileHandler deletes the profile of the user
func DeleteProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the user ID from the context (set by the AuthMiddleware)
	userID := r.Context().Value("userID").(uint)

	// Fetch the user information from the database
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Delete the user from the database
	if err := db.DB.Delete(&user).Error; err != nil {
		http.Error(w, "Failed to delete user profile", http.StatusInternalServerError)
		return
	}

	// Return a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User profile deleted successfully")
}
