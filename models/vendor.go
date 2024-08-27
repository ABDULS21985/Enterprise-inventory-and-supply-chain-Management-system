package models

import "gorm.io/gorm"

type Vendor struct {
	gorm.Model
	Name        string `json:"name"`
	ContactInfo string `json:"contact_info"`
}
