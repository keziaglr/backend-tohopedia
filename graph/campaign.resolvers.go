package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *queryResolver) Campaigns(ctx context.Context) ([]*model.Campaign, error) {
	var campaigns []*model.Campaign
	r.DB.Find(&campaigns)
	return campaigns, nil
}
