package model

import (
	"time"
)

type Shop struct {
	ID                int               `json:"id" gorm:"primaryKey"`
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
	TypeID			int					`json:"type_id"`
	Type			*ShopType			`json:"shop_type" gorm:"foreignKey:TypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BadgesID          int               `json:"badges_id"`
	Badges            *Badges           `json:"badges" gorm:"foreignKey:BadgesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product           []*Product        `json:"product" gorm:"many2many:shop_product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User              *User             `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt         time.Time         `json:"createdAt"`
	UpdatedAt         time.Time         `json:"updatedAt"`
	DeletedAt         time.Time         `json:"deletedAt"`
}

type ShopShippingVendor struct {
	ShopID   int             `json:"shop_id" gorm:"primaryKey"`
	VendorID int             `json:"vendor_id"`
	Shop     *Shop           `json:"shop" gorm:"foreignKey:ShopID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Vendor   *ShippingVendor `json:"vendor" gorm:"foreignKey:VendorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ShopVoucher struct {
	ShopID    int       `json:"shop_id" gorm:"primaryKey"`
	VoucherID int       `json:"voucher_id"`
	Shop      *Shop     `json:"shop" gorm:"foreignKey:ShopID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Voucher   *Voucher  `json:"voucher" gorm:"foreignKey:VoucherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}