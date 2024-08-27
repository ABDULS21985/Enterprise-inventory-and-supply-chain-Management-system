package services

import (
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"

	"golang.org/x/crypto/bcrypt"
)

// CreateProfile creates a new profile with a hashed password
func CreateProfile(profile models.Profile) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(profile.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	profile.Password = string(hashedPassword)

	// Save the profile to the database
	result := db.DB.Create(&profile)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetProfile retrieves a profile by ID
func GetProfile(id uint) (models.Profile, error) {
	var profile models.Profile
	result := db.DB.First(&profile, id)
	if result.Error != nil {
		return profile, result.Error
	}

	return profile, nil
}

// GetProfileByEmail retrieves a profile by email
func GetProfileByEmail(email string) (models.Profile, error) {
	var profile models.Profile
	result := db.DB.Where("email = ?", email).First(&profile)
	if result.Error != nil {
		return profile, result.Error
	}

	return profile, nil
}

// UpdateProfile updates a profile's details
func UpdateProfile(profile models.Profile) error {
	result := db.DB.Save(&profile)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteProfile deletes a profile by ID
func DeleteProfile(id uint) error {
	result := db.DB.Delete(&models.Profile{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ListProfiles retrieves all profiles
func ListProfiles() ([]models.Profile, error) {
	var profiles []models.Profile
	result := db.DB.Find(&profiles)
	if result.Error != nil {
		return profiles, result.Error
	}

	return profiles, nil
}

// UpdatePassword updates a profile's password

func UpdatePassword(profile models.Profile) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(profile.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	profile.Password = string(hashedPassword)

	// Save the updated profile information
	result := db.DB.Save(&profile)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
