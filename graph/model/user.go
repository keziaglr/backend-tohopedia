package model

import (
	"time"
)

type User struct {
	ID              int        `json:"id" gorm:"primaryKey"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Password        string     `json:"password"`
	Dob             string     `json:"dob"`
	Gender          string     `json:"gender"`
	PhoneNumber     string     `json:"phoneNumber"`
	ProfilePicture  string     `json:"profilePicture"`
	IsSuspend       bool       `json:"isSuspend"`
	ShippingAddress string     `json:"shippingAddress"`
	Role            string     `json:"role"`
	Voucher         []*Voucher `json:"voucher" gorm:"many2many:user_voucher;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Wishlist        []*Product `json:"wishlist" gorm:"many2many:user_wishlist;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	DeletedAt       time.Time  `json:"deletedAt"`
}
