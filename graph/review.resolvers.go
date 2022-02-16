package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *mutationResolver) CreateReview(ctx context.Context, userID int, transactionID int, score int, description string, image string, typeReview string) (*model.Review, error) {
	review := model.Review{
		UserID:      userID,
		Score:       score,
		Description: description,
		Image:       image,
		Type:        typeReview,
	}

	r.DB.Create(&review)
	var transaction []*model.TransactionDetail
	r.DB.Where("transaction_id = ?", transactionID).Find(&transaction)
	for i := 0; i < len(transaction); i++ {
		r.DB.Exec("INSERT INTO product_review VALUES (?, ?)", transaction[i].ProductID, review.ID)
	}

	return &review, nil
}

func (r *mutationResolver) CreateReviewReply(ctx context.Context, reviewID int, sourceID int, role string, messsage string) (*model.ReviewReply, error) {
	review := model.ReviewReply{
		ReviewID: reviewID,
		SourceID: sourceID,
		Role:     role,
		Messsage: messsage,
	}

	r.DB.Create(&review)
	return &review, nil
}

func (r *queryResolver) GetReviewsByType(ctx context.Context, productID int, typeReview string, filter string) ([]*model.Review, error) {
	var review []*model.Review
	var temp = r.DB.Select("reviews.*").Table("reviews").Joins("join product_review on product_review.review_id = reviews.id").Where("product_review.product_id = ?", productID)

	if filter != ""{
		if filter == "1" || filter == "2" || filter == "3" || filter == "4" || filter == "5" {
			temp.Where("score = ?", filter)
		} else if filter == "img" {
			temp.Where("image != ?", "null")
		}
	}

	temp.Find(&review)
	return review, nil
}

func (r *queryResolver) GetReviewDetail(ctx context.Context, reviewID int) ([]*model.ReviewReply, error) {
	var review []*model.ReviewReply
	r.DB.Where("review_id = ?", reviewID).Find(&review)
	return review, nil
}
