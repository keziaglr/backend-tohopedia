package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	var categories []*model.Category
	r.DB.Find(&categories)
	return categories, nil
}

func (r *queryResolver) GetSubCategories(ctx context.Context, categoryID int) ([]*model.SubCategory, error) {
	panic(fmt.Errorf("not implemented"))
}
