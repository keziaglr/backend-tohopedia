package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *queryResolver) GetBadge(ctx context.Context, shopID int) (*model.Badges, error) {
	var badge *model.Badges

	var shop *model.Shop
	r.DB.Table("shops").Where("id=?", shopID).Scan(&shop)

	r.DB.Where("? BETWEEN start_point AND end_point", shop.Points).Find(&badge)

	return badge, nil
}
