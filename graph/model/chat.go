package model

import (
	"time"
)

type Chat struct {
	ID        int       `json:"id"`
	ShopID    int       `json:"shop_id"`
	UserID    int       `json:"user_id"`
	Shop      *Shop     `json:"shop" gorm:"foreignKey:ShopID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      *User     `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type ChatDetail struct {
	ID        int       `json:"id"`
	ChatID    int       `json:"chat_id"`
	SourceID  int       `json:"source_id"`
	Chat      *Chat     `json:"chat" gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role      string    `json:"role"`
	Message   string   `json:"message"`
	Image     string   `json:"image"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}