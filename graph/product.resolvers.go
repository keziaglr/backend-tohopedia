package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"
	"time"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *mutationResolver) InsertProduct(ctx context.Context, name string, categoryID int, images []string, description string, price int, discount *int, metaData *string) (*model.Product, error) {
	product := model.Product{
		Name:          name,
		SubCategoryID: categoryID,
		Images:        []*model.ProductImage{{URL: images[1]}, {URL: images[2]}},
		Description:   description,
		Price:         price,
		Discount:      *discount,
		MetaData:      *metaData,
	}

	r.DB.Create(&product)
	return &product, nil
}

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
	r.DB.Select("DISTINCT products.*").Table("products").Joins("join shop_product on shop_product.product_id = products.id").Where("shop_product.shop_id = ?", shopID).Preload("Images").Find(&products)
	return products, nil
}

func (r *queryResolver) GetProductsTopDisc(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	r.DB.Limit(15).Order("discount DESC").Preload("Images").Find(&products)
	return products, nil
}

func (r *queryResolver) GetProductsByCategories(ctx context.Context, categoryID int) ([]*model.Product, error) {
	var products []*model.Product
	r.DB.Select("products.*").Table("products").Joins("join sub_categories on products.sub_category_id = sub_categories.id").Where("sub_categories.category_id = ?", categoryID).Preload("Images").Find(&products)
	return products, nil
}

func (r *queryResolver) GetProductsSearch(ctx context.Context, search string, sort *string, input *model.Filter) ([]*model.Product, error) {
	var products []*model.Product
	var name = "%" + search + "%"
	var temp = r.DB.Select("DISTINCT products.*").Table("shops").Joins("join shop_product on shops.id = shop_product.shop_id").Joins("join products on products.id = shop_product.product_id").Joins("join product_image on products.id = product_image.product_id").Joins("join product_images on product_images.id = product_image.product_image_id").Joins("join shop_shipping_vendors on shop_shipping_vendors.shop_id = shops.id").Joins("join shipping_vendors on shop_shipping_vendors.vendor_id = shipping_vendors.id")

	if input.Type != nil {
		temp = temp.Where("shops.type_id IN ?", input.Type)
	}

	if input.Location != nil {
		temp = temp.Where("shops.address IN ?", input.Location)
	}

	if input.MinPrice != nil {
		temp = temp.Where("products.price >= ?", input.MinPrice)
	}

	if input.MaxPrice != nil {
		temp = temp.Where("products.price <= ?", input.MaxPrice)
	}

	if input.Courier != nil {
		temp = temp.Where("shop_shipping_vendors.vendor_id IN ?", input.Courier)
	}

	if input.Rating != nil {
		temp = temp.Where("products.rating >= ?", input.Rating)
	}

	if input.ShippingTime != nil {
		temp = temp.Where("shipping_vendors.delivery_time = ?", input.ShippingTime)
	}

	if input.ProductAdded != nil {
		temp = temp.Where("DATEDIFF(?, products.created_at) <= ?", time.Now(), input.ProductAdded)
	}

	temp = temp.Where("products.name LIKE ?", name).Preload("Images")
	if strings.Compare(*sort, "suitable") == 0 {
		var by = ""
		if strings.Contains(search, " ") {
			by = "products.name like '$" + search + "%'"
		} else {
			by = "products.name = '$" + search + "'"
		}
		temp = temp.Order(by + " desc")
	} else if strings.Compare(*sort, "rating") == 0 {
		temp = temp.Order("rating desc")
	} else if strings.Compare(*sort, "latest") == 0 {
		temp = temp.Order("created_at desc")
	} else if strings.Compare(*sort, "highPrice") == 0 {
		temp = temp.Order("price desc")
	} else if strings.Compare(*sort, "lowPrice") == 0 {
		temp = temp.Order("price asc")
	}

	temp.Find(&products).Preload("Images")
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
	return products, nil
}

func (r *queryResolver) GetBestSellingProducts(ctx context.Context, shopID int) ([]*model.Product, error) {
	var products []*model.Product
	r.DB.Limit(10).Select("DISTINCT products.*").Table("products").Joins("join shop_product on shop_product.product_id = products.id").Where("shop_product.shop_id = ?", shopID).Preload("Images").Order("sold_count DESC").Find(&products)
	return products, nil
}
