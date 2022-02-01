package graph


				// This file will be automatically regenerated based on the schema, any resolver implementations
				// will be copied through when generating and any unknown code will be moved to the end.

import (
"context"
"fmt"
"io"
"strconv"
"time"
"sync"
"errors"
"bytes"
gqlparser "github.com/vektah/gqlparser/v2"
"github.com/vektah/gqlparser/v2/ast"
"github.com/99designs/gqlgen/graphql"
"github.com/99designs/gqlgen/graphql/introspection"
"github.com/keziaglr/backend-tohopedia/graph/model")


















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
		var products []*model.Product
	db.Joins("ShopType").Joins("Shop").Preload("Product").Preload("ProductImage").Find(&products)
	return products, nil
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

func (r *queryResolver) GetProductsSearch(ctx context.Context, search string, sort *string, input *model.Filter) ([]*model.Product, error) {
		var products []*model.Product
var name = "%" + search + "%"
var temp


var filters *model.Filter
if filters != nil {
	if input.Type != nil {
		temp = db.Joins("ShopType").Joins("Shop").Preload("Product").Preload("ProductImage").Find(&products, "shop_type.id IN ?", input.Type)
	}

} else {
	temp = r.DB.Where("name LIKE ?", name).Preload("Images")
	if strings.Compare(*sort, "suitable") == 0 {
		var by = ""
		if strings.Contains(search, " ") {
			by = "name like '$" + search + "%'"
		} else {
			by = "name = '$" + search + "'"
		}
		temp.Order(by + " desc")
	} else if strings.Compare(*sort, "rating") == 0 {
		temp.Order("rating desc")
	} else if strings.Compare(*sort, "latest") == 0 {
		temp.Order("created_at desc")
	} else if strings.Compare(*sort, "highPrice") == 0 {
		temp.Order("price desc")
	} else if strings.Compare(*sort, "lowPrice") == 0 {
		temp.Order("price asc")
	}
	
}
return products, nil
	}

func (r *queryResolver) GetProductsMatch(ctx context.Context, search string) ([]*model.Product, error) {
		var products []*model.Product
var name = "%" + search + "%"
var by = ""
if strings.Contains(search, " ") {
	by = "name like '$" + search + "%'"
} else {
	by = "name = '$" + search + "'"
}
r.DB.Limit(3).Where("name LIKE ?", name).Order(by + " DESC").Preload("Images").Find(&products)
return products, ni
	}








