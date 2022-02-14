package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *queryResolver) Vendors(ctx context.Context) ([]*model.ShippingVendor, error) {
	var vendors []*model.ShippingVendor

	r.DB.Find(&vendors)
	return vendors, nil
}

func (r *queryResolver) GetVendorByProduct(ctx context.Context, productID int) ([]*model.ShippingVendor, error) {
	var vendors []*model.ShippingVendor

	var shop *model.Shop
	r.DB.Table("shops").Select("shops.*").Joins("left join shop_product on shops.id = shop_product.shop_id").Where("shop_product.product_id=?", productID).Scan(&shop)

	r.DB.Table("shipping_vendors").Select("shipping_vendors.*").Joins("left join shop_shipping_vendors on shipping_vendors.id = shop_shipping_vendors.vendor_id").Where("shop_shipping_vendors.shop_id=?", shop.ID).Scan(&vendors)
	return vendors, nil
}

func (r *queryResolver) GetVendorByUser(ctx context.Context, userID int) ([]*model.ShippingVendor, error) {
	var vendors []*model.ShippingVendor
	var product *model.Product
	r.DB.Select("DISTINCT products.*").Table("products").Joins("join carts on carts.product_id = products.id").Where("user_id = ?", userID).First(&product)

	var shop *model.Shop
	r.DB.Table("shops").Select("shops.*").Joins("join shop_product on shops.id = shop_product.shop_id").Where("shop_product.product_id = ?", product.ID).Scan(&shop)

	r.DB.Table("shipping_vendors").Select("shipping_vendors.*").Joins("left join shop_shipping_vendors on shipping_vendors.id = shop_shipping_vendors.vendor_id").Where("shop_shipping_vendors.shop_id=?", shop.ID).Scan(&vendors)
	return vendors, nil
}
