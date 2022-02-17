// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	OtpCode  string `json:"otpCode"`
}

type CartProduct struct {
	ProductID []int `json:"productId"`
	Qty       []int `json:"qty"`
}

type CreateShop struct {
	UserID      int    `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	NameSlug    string `json:"nameSlug"`
	Address     string `json:"address"`
}

type CreateVoucher struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	DiscountRate int    `json:"discountRate"`
	Tnc          string `json:"tnc"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
}

type Filter struct {
	Type         []*int    `json:"type"`
	Location     []*string `json:"location"`
	MinPrice     *int      `json:"minPrice"`
	MaxPrice     *int      `json:"maxPrice"`
	Courier      []*int    `json:"courier"`
	Rating       *int      `json:"rating"`
	ShippingTime *int      `json:"shippingTime"`
	ProductAdded *int      `json:"productAdded"`
}

type FilterTransaction struct {
	Keyword *string `json:"keyword"`
	Status  *string `json:"status"`
	Date    *string `json:"date"`
}

type InsertMetaData struct {
	Label []*string `json:"label"`
	Value []*string `json:"value"`
}

type MetaData struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}

type ShippingAddress struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
}

type TopUp struct {
	ID    int    `json:"id"`
	Value int    `json:"value"`
	Code  string `json:"code"`
}

type UpdateShop struct {
	ProfilePicture    string `json:"profilePicture"`
	Name              string `json:"name"`
	NameSlug          string `json:"nameSlug"`
	Slogan            string `json:"slogan"`
	Description       string `json:"description"`
	OperationalHour   string `json:"operationalHour"`
	OperationalStatus string `json:"operationalStatus"`
}

type UpdateUser struct {
	ProfilePicture string   `json:"profilePicture"`
	Name           string   `json:"name"`
	Dob            string   `json:"dob"`
	Gender         string   `json:"gender"`
	Email          string   `json:"email"`
	PhoneNumber    string   `json:"phoneNumber"`
	Address        []string `json:"address"`
}
