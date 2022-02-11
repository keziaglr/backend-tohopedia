package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/keziaglr/backend-tohopedia/graph/model"
	"gorm.io/gorm"
)

func (r *mutationResolver) CreateCart(ctx context.Context, userID int, productID int, qty int, note string) (*model.Cart, error) {
	var cart *model.Cart
	err := r.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if qty > 0 {
			cart1 := model.Cart{
				UserID:    userID,
				ProductID: productID,
				Qty:       qty,
				Note:      note,
			}
			r.DB.Create(&cart1)
			return &cart1, nil
		}
	} else if cart != nil {
		var cartQty = cart.Qty
		cartQty += qty
		if cartQty <= 0 {
			r.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&cart)
		} else {
			cart.Qty = cartQty
			r.DB.Save(&cart)
		}
		return cart, nil
	}

	return nil, nil
}

func (r *mutationResolver) DeleteCart(ctx context.Context, userID int, productID int) (*model.Cart, error) {
	var cart *model.Cart
	r.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart)

	if cart != nil {
		r.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&cart)
		return cart, nil
	}

	return nil, nil
}

func (r *mutationResolver) Checkout(ctx context.Context, userID int, productID []int) (string, error) {
	var cart *model.Cart
	r.DB.Where("user_id = ? AND product_id IN ?", userID, productID).Delete(&cart)
	return "Success Delete", nil

	return "Not Found", nil
}

func (r *queryResolver) Carts(ctx context.Context, userID int) ([]*model.Product, error) {
	var products []*model.Product
	r.DB.Select("DISTINCT products.*").Table("products").Joins("join carts on carts.product_id = products.id").Where("user_id = ?", userID).Preload("Images").Find(&products)
	return products, nil
}

func (r *queryResolver) Carts2(ctx context.Context, userID int) ([]*model.Cart, error) {
	var carts []*model.Cart
	r.DB.Select("DISTINCT products.*").Table("products").Joins("join carts on carts.product_id = products.id").Where("user_id = ?", userID).Preload("Images").Find(&carts)
	return carts, nil
}
