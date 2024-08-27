package controllers

import (
	"encoding/json"
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
	"inventory-supply-chain-system/pkg/utils"
	"log"
	"net/http"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser registers a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Check if user with the given email already exists
	var existingUser models.User
	if db.DB.Where("email = ?", user.Email).First(&existingUser).Error == nil {
		http.Error(w, "User with this email already exists", http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Format the permissions as a PostgreSQL array
	user.Permissions = pq.StringArray([]string{"view_reports", "approve_transactions", "manage_team"})

	// Save the user to the database
	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	// Debugging log: output the hashed password
	log.Printf("User registered: %s, Hashed Password: %s", user.Email, user.Password)

	// Return the created user (without the password)
	user.Password = ""
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// LoginUser authenticates the user and returns a JWT token
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Find the user by email
	var user models.User
	if err := db.DB.Where("email = ?", loginUser.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Debugging log: output the stored hashed password
	log.Printf("Login attempt for: %s, Stored Hashed Password: %s", user.Email, user.Password)

	// Compare the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		log.Printf("Password comparison failed for: %s, Provided Password: %s", user.Email, loginUser.Password)
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Return the token in the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
