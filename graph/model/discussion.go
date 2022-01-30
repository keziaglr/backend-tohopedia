package model

import (
	"time"
)

type Discussion struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	ProductID   int       `json:"product_id"`
	Content     string    `json:"content"`
	Description string    `json:"description"`
	User        *User     `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product     *Product  `json:"product" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

type DiscussionReply struct {
	ID           int         `json:"id"`
	DiscussionID int         `json:"discussion_id"`
	SourceID     int         `json:"source_id"`
	Role         string      `json:"role"`
	Messsage     string      `json:"messsage"`
	Discussion   *Discussion `json:"discussion" gorm:"foreignKey:DiscussionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
	DeletedAt    time.Time   `json:"deletedAt"`
}
