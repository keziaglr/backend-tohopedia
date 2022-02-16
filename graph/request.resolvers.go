package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *mutationResolver) UpdateStatusUser(ctx context.Context, userID int, status bool) (*model.User, error) {
	var user *model.User
	r.DB.Where("id = ?", userID).First(&user)

	if user != nil {
		user.IsSuspend = status
		r.DB.Save(&user)
		return user, nil
	}
	return nil, nil
}

func (r *mutationResolver) SendRequest(ctx context.Context, userID int, status string) (*model.Request, error) {
	req := model.Request{
		UserID: userID,
		Status: "Pending",
	}

	r.DB.Create(&req)
	return &req, nil
}

func (r *mutationResolver) ResponseRequest(ctx context.Context, userID int, status bool, requestID int) (*model.Request, error) {
	var user *model.User
	r.DB.Where("id = ?", userID).First(&user)

	if user != nil {
		user.IsSuspend = status
		r.DB.Save(&user)

		var req *model.Request
		r.DB.Where("id = ?", requestID).First(&req)
		if status == true {
			req.Status = "Declined"
		} else if status == false {
			req.Status = "Approved"
		}
		r.DB.Save(req)
	}
	return nil, nil
}

func (r *queryResolver) Requests(ctx context.Context) ([]*model.Request, error) {
	var requests []*model.Request
	r.DB.Find(&requests)
	return requests, nil
}
