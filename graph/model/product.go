package model

import (
	"time"
)

type Product struct {
	ID            int             `json:"id" gorm:"primaryKey"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	Price         int             `json:"price"`
	Discount      int             `json:"discount"`
	MetaData      string          `json:"metaData"`
	AddedTime     time.Time       `json:"addedTime"`
	Stock         int             `json:"stock"`
	Rating        int             `json:"rating"`
	SoldCount	int 			`json:"sold_count"`
	SubCategoryID int             `json:"sub_category_id"`
	SubCategory   *SubCategory    `json:"sub_category" gorm:"foreignKey:SubCategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Images        []*ProductImage `json:"images" gorm:"many2many:product_image;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	DeletedAt     time.Time       `json:"deletedAt"`
}

type ProductImage struct {
	ID  int    `json:"id" gorm:"primaryKey"`
	URL string `json:"url"`
}