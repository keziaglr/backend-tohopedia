package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *mutationResolver) CreateShop(ctx context.Context, input model.CreateShop) (*model.Shop, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateShop(ctx context.Context, id int, input model.UpdateShop) (*model.Shop, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetShopByProduct(ctx context.Context, productID int) (*model.Shop, error) {
	var shop *model.Shop
	r.DB.Preload("Product", "id=?", productID).First(&shop)
	return shop, nil
}
