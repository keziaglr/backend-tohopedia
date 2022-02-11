package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *mutationResolver) CreateShop(ctx context.Context, input model.CreateShop) (*model.Shop, error) {
	shop := model.Shop{
		UserID:      input.UserID,
		PhoneNumber: input.PhoneNumber,
		Name:        input.Name,
		NameSlug:    input.NameSlug,
		Address:     input.Address,
		BadgesID:    1,
		TypeID:      3,
		Points:      0,
	}

	r.DB.Create(&shop)
	return &shop, nil
}

func (r *mutationResolver) UpdateShop(ctx context.Context, id int, input model.UpdateShop) (*model.Shop, error) {
	var shop *model.Shop
	r.DB.Where("id = ?", id).First(&shop)

	if shop != nil {
		shop.Image = input.ProfilePicture
		shop.Name = input.Name
		shop.NameSlug = input.NameSlug
		shop.Slogan = input.Slogan
		shop.Description = input.Description
		shop.OperationalHour = input.OperationalHour
		shop.OperationalStatus = input.OperationalStatus
		r.DB.Save(&shop)
		return shop, nil
	}
	return nil, nil
}

func (r *queryResolver) GetShopByProduct(ctx context.Context, productID int) (*model.Shop, error) {
	var shop *model.Shop
	r.DB.Table("shops").Select("shops.*").Joins("join shop_product on shops.id = shop_product.shop_id").Where("shop_product.product_id=?", productID).Scan(&shop)
	return shop, nil
}

func (r *queryResolver) GetShopMatch(ctx context.Context, search string) (*model.Shop, error) {
	var shop *model.Shop
	r.DB.Select("shops.*").Table("products").Joins("join shop_product on products.id = shop_product.product_id").Joins("join shops on shops.id = shop_product.shop_id").Where("products.name LIKE ?", "%"+search+"%").Group("shops.name").Order("COUNT(DISTINCT products.id) desc").Scan(&shop)
	return shop, nil
}

func (r *queryResolver) GetShopByID(ctx context.Context, shopID int) (*model.Shop, error) {
	var shop *model.Shop
	r.DB.Where("id = ?", shopID).Preload("Promo").First(&shop)
	return shop, nil
}

func (r *queryResolver) GetPromoByShop(ctx context.Context, shopID int) ([]*model.ShopPromo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetShopByUser(ctx context.Context, userID int) (*model.Shop, error) {
	var shop *model.Shop
	r.DB.Where("user_id = ?", userID).First(&shop)
	return shop, nil
}
