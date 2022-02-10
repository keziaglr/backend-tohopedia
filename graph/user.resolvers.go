package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/keziaglr/backend-tohopedia/graph/generated"
	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.AuthUser) (*model.User, error) {
	var otp model.Otp
	r.DB.Debug().First(&otp, "code=?", input.OtpCode)

	if otp.ID == 0 {
		return nil, nil
	}

	if time.Since(otp.ValidTime).Minutes() >= 2 {
		r.DB.Delete(&otp, "code=?", input.OtpCode)
		return nil, nil
	}

	r.DB.Delete(&otp, "code=?", input.OtpCode)

	user := model.User{
		Email:    input.Email,
		Password: input.Password,
	}

	r.DB.Create(&user)
	return &user, nil
}

func (r *mutationResolver) AuthUser(ctx context.Context, input model.AuthUser) (*model.User, error) {
	var user *model.User
	r.DB.Where("email = ? AND password = ?", input.Email, input.Password).First(&user)

	if user == nil {
		return nil, nil
	}

	var otp model.Otp
	r.DB.Debug().First(&otp, "code=?", input.OtpCode)

	if otp.ID == 0 {
		fmt.Printf("OTP ID = 0")
		return nil, nil
	}

	if time.Since(otp.ValidTime).Minutes() >= 2 {
		r.DB.Delete(&otp, "code=?", input.OtpCode)
		return nil, nil
	}

	if user == nil {
		fmt.Printf("User nil")
	}
	r.DB.Delete(&otp, "code=?", input.OtpCode)
	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input model.UpdateUser) (*model.User, error) {
	var user *model.User
	r.DB.Where("id = ?", id).First(&user)

	if user != nil {
		user.ProfilePicture = input.ProfilePicture
		user.Name = input.Name
		user.Dob = input.Dob
		user.Gender = input.Gender
		user.Email = input.Email
		user.PhoneNumber = input.PhoneNumber
		r.DB.Save(&user)
		return user, nil
	}
	return nil, nil
}

func (r *mutationResolver) ResetPassword(ctx context.Context, input model.AuthUser) (*model.User, error) {
	var user *model.User
	r.DB.Where("email = ?", input.Email).First(&user)

	if user == nil {
		return nil, nil
	}

	var otp model.Otp
	r.DB.Debug().First(&otp, "code=?", input.OtpCode)

	if otp.ID == 0 {
		return nil, nil
	}

	if time.Since(otp.ValidTime).Minutes() >= 2 {
		r.DB.Delete(&otp, "code=?", input.OtpCode)
		return nil, nil
	}

	r.DB.Delete(&otp, "code=?", input.OtpCode)

	if user != nil {
		user.Password = input.Password
		r.DB.Save(&user)
		return user, nil
	}
	return nil, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	r.DB.Where("role = ?", "user").Find(&users)
	return users, nil
}

func (r *queryResolver) GetUserByEmailPass(ctx context.Context, email string, password string) (*model.User, error) {
	var user *model.User
	r.DB.Where("email = ? AND password = ?", email, password).First(&user)

	if user == nil {
		return nil, nil
	}
	return user, nil
}

func (r *queryResolver) GetUserAuth(ctx context.Context, input model.AuthUser) (*model.User, error) {
	var user *model.User
	r.DB.Where("email = ? AND password = ?", input.Email, input.Password).First(&user)

	return user, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	var user *model.User
	r.DB.Where("id = ?", id).First(&user)

	return user, nil
}

func (r *queryResolver) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user *model.User
	r.DB.Where("email = ?", email).First(&user)

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
