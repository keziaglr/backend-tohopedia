package model

import (
	"time"
)

type Shop struct {
	ID                int               `json:"id"`
	UserID            int               `json:"user_id"`
	Name              string            `json:"name"`
	NameSlug          string            `json:"nameSlug"`
	Points            int               `json:"points"`
	Image             string            `json:"image"`
	OperationalStatus string            `json:"operationalStatus"`
	OperationalHour   string            `json:"operationalHour"`
	Description       string            `json:"description"`
	Slogan            string            `json:"slogan"`
	Address           string            `json:"address"`
	PhoneNumber       string            `json:"phoneNumber"`
	BadgesID          int               `json:"badges_id"`
	Badges            *Badges           `json:"badges" gorm:"foreignKey:BadgesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product           []*Product        `json:"product" gorm:"many2many:shop_product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User              *User             `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt         time.Time         `json:"createdAt"`
	UpdatedAt         time.Time         `json:"updatedAt"`
	DeletedAt         time.Time         `json:"deletedAt"`
}

type ShopShippingVendor struct {
	ShopID   int             `json:"shop_id"`
	VendorID int             `json:"vendor_id"`
	Shop     *Shop           `json:"shop" gorm:"foreignKey:ShopID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Vendor   *ShippingVendor `json:"vendor" gorm:"foreignKey:VendorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}