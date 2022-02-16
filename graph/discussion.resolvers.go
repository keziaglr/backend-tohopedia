package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *mutationResolver) CreateDiscussion(ctx context.Context, userID int, productID int, content string) (*model.Discussion, error) {
	discussion := model.Discussion{
		UserID:    userID,
		ProductID: productID,
		Content:   content,
	}

	r.DB.Create(&discussion)
	return &discussion, nil
}

func (r *mutationResolver) CreateDiscussionReply(ctx context.Context, discussionID int, sourceID int, role string, messsage string) (*model.DiscussionReply, error) {
	discussion := model.DiscussionReply{
		DiscussionID: discussionID,
		SourceID:     sourceID,
		Role:         role,
		Messsage:     messsage,
	}

	r.DB.Create(&discussion)
	return &discussion, nil
}

func (r *queryResolver) GetDiscussion(ctx context.Context, productID int) ([]*model.Discussion, error) {
	var discussion []*model.Discussion
	r.DB.Where("product_id=?", productID).Find(&discussion)
	return discussion, nil
}

func (r *queryResolver) GetDiscussionDetail(ctx context.Context, discussionID int) ([]*model.DiscussionReply, error) {
	var discussion []*model.DiscussionReply
	r.DB.Where("discussion_id=?", discussionID).Find(&discussion)
	return discussion, nil
}
