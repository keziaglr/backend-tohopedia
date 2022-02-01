package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *queryResolver) GetVoucherByProduct(ctx context.Context, productID int) ([]*model.Voucher, error) {
	var vouchers []*model.Voucher

	var shop *model.Shop
	r.DB.Table("shops").Select("shops.*").Joins("left join shop_product on shops.id = shop_product.shop_id").Where("shop_product.product_id=?", productID).Scan(&shop)

	r.DB.Table("vouchers").Select("vouchers.*").Joins("left join shop_vouchers on vouchers.id = shop_vouchers.voucher_id").Where("shop_vouchers.shop_id=?", shop.ID).Scan(&vouchers)
	return vouchers, nil
}
