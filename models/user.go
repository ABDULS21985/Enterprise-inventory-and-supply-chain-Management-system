package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Address represents a user's address
type Address struct {
	gorm.Model
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	UserID  uint   `json:"user_id"` // This will be my foreign key to the User model
}

// User represents the user model

type User struct {
	gorm.Model
	Name        string         `json:"name"`
	Email       string         `json:"email" gorm:"unique"`
	Password    string         `json:"-"`
	Role        string         `json:"role"`
	Verified    bool           `json:"verified"`
	Permissions pq.StringArray `json:"permissions" gorm:"type:text[]"`
	Phone       string         `json:"phone"`
	Addresses   []Address      `json:"addresses"`
}
