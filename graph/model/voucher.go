package model

type Voucher struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	DiscountRate int       `json:"discountRate"`
	Tnc          string    `json:"tnc"`
	Code 		string 		`json:"code"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
}