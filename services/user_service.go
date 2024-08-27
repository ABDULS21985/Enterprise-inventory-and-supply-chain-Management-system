package services

import (
	"inventory-supply-chain-system/db"
	"inventory-supply-chain-system/models"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a new user with a hashed password
func CreateUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if result := db.DB.Create(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

// GetUser retrieves a user by ID
func GetUser(id uint) (models.User, error) {
	var user models.User
	if result := db.DB.First(&user, id); result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if result := db.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// UpdateUser updates a user's details
func UpdateUser(user models.User) error {
	if result := db.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id uint) error {
	if result := db.DB.Delete(&models.User{}, id); result.Error != nil {
		return result.Error
	}

	return nil
}

// ListUsers retrieves all users
func ListUsers() ([]models.User, error) {
	var users []models.User
	if result := db.DB.Find(&users); result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

// LoginUser verifies a user's credentials and logs them in
func LoginUser(email, password string) (models.User, error) {
	var user models.User
	if result := db.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return user, result.Error
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, err
	}

	return user, nil
}

// ChangePassword changes a user's password
func ChangePassword(id uint, password string) error {
	var user models.User
	if result := db.DB.First(&user, id); result.Error != nil {
		return result.Error
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if result := db.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

// ResetPassword resets a user's password by email
func ResetPassword(email, password string) error {
	var user models.User
	if result := db.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return result.Error
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if result := db.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

// VerifyUser verifies a user's email
func VerifyUser(email string) error {
	var user models.User
	if result := db.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return result.Error
	}

	user.Verified = true
	if result := db.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

// UnverifyUser unverifies a user's email
func UnverifyUser(email string) error {
	var user models.User
	if result := db.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return result.Error
	}

	user.Verified = false
	if result := db.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

// AddRole adds a role to a user
func AddRole(id uint, role string) error {
	var user models.User
	if result := db.DB.First(&user, id); result.Error != nil {
		return result.Error
	}

	user.Role = role
	if result := db.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

// RemoveRole removes a role from a user
func RemoveRole(id uint) error {
	var user models.User
	if result := db.DB.First(&user, id); result.Error != nil {
		return result.Error
	}

	user.Role = ""
	if result := db.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

// AddPermission adds a permission to a user
func AddPermission(id uint, permission string) error {
	var user models.User
	if result := db.DB.First(&user, id); result.Error != nil {
		return result.Error
	}

	user.Permissions = append(user.Permissions, permission)
	if result := db.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

// RemovePermission removes a permission from a user
func RemovePermission(id uint, permission string) error {
	var user models.User
	if result := db.DB.First(&user, id); result.Error != nil {
		return result.Error
	}

	for i, p := range user.Permissions {
		if p == permission {
			user.Permissions = append(user.Permissions[:i], user.Permissions[i+1:]...)
			break
		}
	}

	if result := db.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

// AddAddress adds an address to a user
func AddAddress(id uint, address models.Address) error {
	var user models.User
	if result := db.DB.First(&user, id); result.Error != nil {
		return result.Error
	}

	user.Addresses = append(user.Addresses, address)
	if result := db.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

// RemoveAddress removes an address from a user
func RemoveAddress(id uint, address models.Address) error {
	var user models.User
	if result := db.DB.First(&user, id); result.Error != nil {
		return result.Error
	}

	for i, a := range user.Addresses {
		if a == address {
			user.Addresses = append(user.Addresses[:i], user.Addresses[i+1:]...)
			break
		}
	}

	if result := db.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
