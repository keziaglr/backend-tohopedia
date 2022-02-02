package model

import (
	"time"
)

type Voucher struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	DiscountRate int       `json:"discountRate"`
	Tnc          string    `json:"tnc"`
	Code 		string 		`json:"code"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
}