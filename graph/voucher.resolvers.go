package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *mutationResolver) CreateUserVoucher(ctx context.Context, voucherID int, userID int) (*model.UserVoucher, error) {
	userVoucher := model.UserVoucher{
		VoucherID: voucherID,
		UserID:    userID,
	}

	r.DB.Create(&userVoucher)
	return &userVoucher, nil
}

func (r *queryResolver) GetVoucherByProduct(ctx context.Context, productID int) ([]*model.Voucher, error) {
	var vouchers []*model.Voucher

	var shop *model.Shop
	r.DB.Table("shops").Select("shops.*").Joins("left join shop_product on shops.id = shop_product.shop_id").Where("shop_product.product_id=?", productID).Scan(&shop)

	r.DB.Table("vouchers").Select("vouchers.*").Joins("left join shop_vouchers on vouchers.id = shop_vouchers.voucher_id").Where("shop_vouchers.shop_id=?", shop.ID).Scan(&vouchers)
	return vouchers, nil
}

func (r *queryResolver) Vouchers(ctx context.Context) ([]*model.Voucher, error) {
	var vouchers []*model.Voucher
	r.DB.Find(&vouchers)
	return vouchers, nil
}

func (r *queryResolver) GetVoucherByID(ctx context.Context, voucherID int) (*model.Voucher, error) {
	var voucher *model.Voucher
	r.DB.Where("id = ?", voucherID).Find(&voucher)
	return voucher, nil
}

func (r *queryResolver) GetVoucherCart(ctx context.Context, userID int) ([]*model.Voucher, error) {
	var voucher []*model.Voucher
	r.DB.Table("vouchers").Select("vouchers.*").Joins("join user_vouchers on vouchers.id = user_vouchers.voucher_id").Where("user_vouchers.user_id=?", userID).Scan(&voucher)
	return voucher, nil
}
