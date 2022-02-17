package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *mutationResolver) Topup(ctx context.Context, code string, value int, userID int) (*model.TopUp, error) {
	var topup *model.TopUp
	r.DB.Where("value = ? AND code = ?", value, code).Find(&topup)

	if topup != nil {
		var user *model.User
		r.DB.Where("id=?", userID).Find(&user)

		if user != nil {
			user.Balance += value
			r.DB.Save(&user)
			return topup, nil
		}
	}

	return nil, nil
}
