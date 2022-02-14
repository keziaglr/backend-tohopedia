package model

import (
	"time"
)

type TransactionDetail struct {
	ID            int                `json:"id"`
	TransactionID int                `json:"transaction_id"`
	ProductID     int                `json:"product_id"`
	Qty           int                `json:"qty"`
	Transaction   *TransactionHeader `json:"transaction" gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product       *Product           `json:"product" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	DeletedAt     time.Time          `json:"deletedAt"`
}

type TransactionHeader struct {
	ID              int             `json:"id" gorm:"primaryKey"`
	UserID          int             `json:"user_id"`
	TransactionType string          `json:"transactionType"`
	TransactionDate time.Time          `json:"transactionDate"`
	Voucher       *Voucher           `json:"voucher" gorm:"foreignKey:VoucherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VoucherID     int                `json:"voucher_id"`
	Status          string          `json:"status"`
	InvoiceNumber   string          `json:"invoiceNumber"`
	PaymentMethod   string          `json:"paymentMethod"`
	ShippingAddress string          `json:"shippingAddress"`
	PaymentDiscount int             `json:"paymentDiscount"`
	ShippingID      int             `json:"shipping_id"`
	ShippingVendor  *ShippingVendor `json:"shippingVendor" gorm:"foreignKey:ShippingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User            *User           `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
	DeletedAt       time.Time       `json:"deletedAt"`
}
