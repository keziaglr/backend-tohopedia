// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type RegisterUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	ProfilePicture string `json:"profilePicture"`
	Name           string `json:"name"`
	Dob            string `json:"dob"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
}

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Dob            string `json:"dob"`
	Gender         string `json:"gender"`
	PhoneNumber    string `json:"phoneNumber"`
	ProfilePicture string `json:"profilePicture"`
	IsSuspend      bool   `json:"isSuspend"`
}
