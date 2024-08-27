package models

import "gorm.io/gorm"

// proile represents the profile model
type Profile struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	User     User   `gorm:"foreignKey:UserID"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
