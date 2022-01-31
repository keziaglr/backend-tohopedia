package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	r.DB.Find(&products)
	return products, nil
}

func (r *queryResolver) GetProductByID(ctx context.Context, id int) (*model.Product, error) {
	var product *model.Product
	r.DB.Preload("Images").Where("id=?", id).First(&product)
	return product, nil
}

func (r *queryResolver) GetProductsByShop(ctx context.Context, shopID int) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProductsTopDisc(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	r.DB.Limit(15).Order("discount DESC").Preload("Images").Find(&products)
	return products, nil
}

func (r *queryResolver) GetProductsByCategories(ctx context.Context, categoryID int) ([]*model.Product, error) {
	var products []*model.Product
	r.DB.Where("sub_category_id=?", categoryID).Find(&products)
	return products, nil
}

func (r *queryResolver) GetProductsSearch(ctx context.Context, search string) ([]*model.Product, error) {
	var products []*model.Product
	r.DB.Where("name LIKE ?", "%name%").Find(&products)
	return products, nil
}

func (r *queryResolver) GetProductOrder(ctx context.Context, by string, order string) ([]*model.Product, error) {
	var products []*model.Product
	r.DB.Order(by + " " + order).Find(&products)
	return products, nil
}

func (r *queryResolver) GetProductFilter(ctx context.Context, by string, value string) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}
