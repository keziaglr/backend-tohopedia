package model

import (
	"time"
)


type Review struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	ProductID   int       `json:"product_id"`
	Score       int       `json:"score"`
	Description string    `json:"description"`
	Image       *string   `json:"image"`
	Status      string    `json:"status"`
	User        *User     `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product     *Product  `json:"product" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

type ReviewReply struct {
	ID        int       `json:"id"`
	ReviewID  int       `json:"review_id"`
	SourceID  int       `json:"source_id"`
	Role      string    `json:"role"`
	Messsage  string    `json:"messsage"`
	Review    *Review   `json:"review" gorm:"foreignKey:ReviewID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}