package models

import "gorm.io/gorm"

type Shipment struct {
	gorm.Model
	OrderID        uint   `json:"order_id"`
	TrackingNumber string `json:"tracking_number"`
	Carrier        string `json:"carrier"`
	ShippingStatus string `json:"shipping_status"`
}
