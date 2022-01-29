package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/generated"
	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.RegisterUser) (*model.User, error) {
	user := model.User{
		Email:    input.Email,
		Password: input.Password,
	}

	r.DB.Create(&user)
	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input model.UpdateUser) (*model.User, error) {
	updatedUser := model.User{
		ProfilePicture: input.ProfilePicture,
		Name:           input.Name,
		Dob:            input.Dob,
		Gender:         input.Gender,
		Email:          input.Gender,
		PhoneNumber:    input.PhoneNumber,
	}
	r.DB.Save(&updatedUser)
	return &updatedUser, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	r.DB.Find(&users)
	return users, nil
}

func (r *queryResolver) GetUserAuth(ctx context.Context, email string, password string) (*model.User, error) {
	var user *model.User
	r.DB.Where("email = ? AND password = ?", email, password).First(&user)
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
