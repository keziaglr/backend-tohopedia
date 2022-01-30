package model

import (
	"time"
)


type Badges struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	StartPoint int    `json:"startPoint"`
	EndPoint   int    `json:"endPoint"`
	Badge      string `json:"badge"`
}


type ShippingVendor struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	DeliveryTime int    `json:"deliveryTime"`
	Price        int    `json:"price"`
}

type Otp struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Code      string    `json:"code"`
	Email     string    `json:"email"`
	ValidTime time.Time `json:"valid_time"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}