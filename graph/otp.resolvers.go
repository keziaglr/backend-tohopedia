package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/keziaglr/backend-tohopedia/graph/model"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

func (r *mutationResolver) CreateOtp(ctx context.Context, email string) (string, error) {
	otp := model.Otp{
		Code:      StringRandom(5),
		Email:     email,
		ValidTime: time.Now(),
	}
	r.DB.Create(&otp)
	// SendOTP(email, otp.Code)
	return otp.Code, nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
const charset = "abcdeABCDE12345"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
func StringRandom(length int) string {
	return StringWithCharset(length, charset)
}
func SendOTP(email, code string) {
	mailjetClient := mailjet.NewMailjetClient("b1e9abc0c2102c71f83a1a2862d23d4a", "39425cec3cc431c722d9043b7eadb5f1")

	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "kezia.lie@binus.ac.id",
				Name:  "Tohopedia Admin",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
					Name:  email,
				},
			},
			Subject:  "Welcome to Tohopedia",
			TextPart: "Welcome New User",
			HTMLPart: "<h3>Dear New User , welcome to <a href='https://www.tokopedia.com/'>Tohopedia</a>!</h3><br/> <h1>This is your OTP : " + code + "</h1>",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
