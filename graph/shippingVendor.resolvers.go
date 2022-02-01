package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *queryResolver) GetVendorByProduct(ctx context.Context, productID int) ([]*model.ShippingVendor, error) {
	var vendors []*model.ShippingVendor

	var shop *model.Shop
	r.DB.Table("shops").Select("shops.*").Joins("left join shop_product on shops.id = shop_product.shop_id").Where("shop_product.product_id=?", productID).Scan(&shop)

	r.DB.Table("shipping_vendors").Select("shipping_vendors.*").Joins("left join shop_shipping_vendors on shipping_vendors.id = shop_shipping_vendors.vendor_id").Where("shop_shipping_vendors.shop_id=?", shop.ID).Scan(&vendors)
	return vendors, nil
}
