package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *queryResolver) GetVendorByShop(ctx context.Context, shopID int) ([]*model.ShippingVendor, error) {
	var vendors []*model.ShippingVendor
	r.DB.Table("shipping_vendors").Select("shipping_vendors.*").Joins("left join shop_shipping_vendors on shipping_vendors.id = shop_shipping_vendors.vendor_id").Where("shop_shipping_vendors.shop_id=?", shopID).Scan(&vendors)
	return vendors, nil
}
