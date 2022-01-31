package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *queryResolver) GetVoucherByShop(ctx context.Context, shopID int) ([]*model.Voucher, error) {
	var vouchers []*model.Voucher
	r.DB.Table("vouchers").Select("vouchers.*").Joins("left join shop_vouchers on vouchers.id = shop_vouchers.voucher_id").Where("shop_vouchers.shop_id=?", shopID).Scan(&vouchers)
	return vouchers, nil
}
