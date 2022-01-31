package model

import (
	"time"
)

type User struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	Dob             string    `json:"dob"`
	Gender          string    `json:"gender"`
	PhoneNumber     string    `json:"phoneNumber"`
	ProfilePicture  string    `json:"profilePicture"`
	IsSuspend       bool      `json:"isSuspend"`
	ShippingAddress string    `json:"shippingAddress"`
	Role            string    `json:"role"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	DeletedAt       time.Time `json:"deletedAt"`
}

type UserVoucher struct {
	VoucherID int       `json:"voucher_id"`
	Voucher   *Voucher  `json:"voucher" gorm:"foreignKey:VoucherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    int       `json:"user_id"`
	User      *User     `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type UserWishlist struct {
	ProductID int       `json:"product_id"`
	Product   *Product  `json:"product" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    int       `json:"user_id"`
	User      *User     `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}