package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *queryResolver) GetTransactionByUser(ctx context.Context, userID int, input *model.FilterTransaction) ([]*model.TransactionHeader, error) {
	var header []*model.TransactionHeader

	var temp = r.DB.Table("transaction_headers").Select("DISTINCT transaction_headers.*").Joins("join transaction_details on transaction_headers.id = transaction_details.transaction_id").Joins("join products on transaction_details.product_id = products.id").Where("transaction_headers.user_id=?", userID)

	if input.Keyword != nil {
		var name = "%" + *input.Keyword + "%"
		temp.Where("products.name LIKE ?", name)
	}
	if input.Status != nil {
		temp.Where("transaction_headers.status = ?", input.Status)
	}
	if input.Date != nil {
		temp.Where("transaction_headers.transaction_date = ?", input.Date)
	}

	temp.Scan(&header)

	return header, nil
}

func (r *queryResolver) GetTransactionDetail(ctx context.Context, userID int, transactionID int) ([]*model.TransactionDetail, error) {
	var detail []*model.TransactionDetail
	r.DB.Table("transaction_details").Select("DISTINCT transaction_details.*").Joins("join transaction_headers on transaction_headers.id = transaction_details.transaction_id").Where("transaction_headers.user_id=? AND transaction_details.transaction_id=?", userID, transactionID).Scan(&detail)
	return detail, nil
}

func (r *queryResolver) GetTransactionByID(ctx context.Context, userID int, id int) (*model.TransactionHeader, error) {
	var header *model.TransactionHeader

	r.DB.Table("transaction_headers").Select("DISTINCT transaction_headers.*").Joins("join transaction_details on transaction_headers.id = transaction_details.transaction_id").Joins("join products on transaction_details.product_id = products.id").Where("transaction_headers.user_id=? AND transaction_headers.id=?", userID, id).First(&header)

	return header, nil
}
