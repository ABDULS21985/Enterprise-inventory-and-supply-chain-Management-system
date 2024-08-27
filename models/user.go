package models

import (
	"gorm.io/gorm"
)

// Address represents a user's address
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

// User represents the user model
type User struct {
	gorm.Model
	Name        string    `json:"name"`
	Email       string    `json:"email" gorm:"unique"`
	Password    string    `json:"-"`
	Role        string    `json:"role"`
	Verified    bool      `json:"verified"`
	Permissions []string  `json:"permissions" gorm:"type:text[]"`
	Addresses   []Address `json:"addresses" gorm:"embedded"`
}
