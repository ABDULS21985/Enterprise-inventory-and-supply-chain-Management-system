package controllers

import (
	"encoding/json"
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"
	"inventory-supply-chain-system/pkg/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser registers a new user
// @Summary Register a new user
// @Description Registers a new user in the system with a hashed password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 400 {string} string "Invalid input"
// @Failure 409 {string} string "User with this email already exists"
// @Failure 500 {string} string "Error hashing password" "Error saving user"
// @Router /auth/register [post]
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

	// Save the user to the database
	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	// Return the created user (without the password)
	user.Password = ""
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// LoginUser authenticates the user and returns a JWT token
// @Summary Login a user
// @Description Authenticates a user and returns a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "User credentials"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "User not found"
// @Failure 401 {string} string "Invalid password"
// @Failure 500 {string} string "Error generating token"
// @Router /auth/login [post]
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

	// Compare the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
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
