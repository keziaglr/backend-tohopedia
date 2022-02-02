package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *queryResolver) GetBadge(ctx context.Context, poin int) (*model.Badges, error) {
	var badge *model.Badges
	r.DB.Where("? BETWEEN start_point AND end_point", poin).Find(&badge)
	return badge, nil
}
