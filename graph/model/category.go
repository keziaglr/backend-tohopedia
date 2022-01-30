package model

type Category struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type SubCategory struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	CategoryID int       `json:"category_id"`
	Name       string    `json:"name"`
	Category   *Category `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}