package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keziaglr/backend-tohopedia/graph/model"
)

func (r *mutationResolver) CreateHeaderChat(ctx context.Context, userID int, shopID int) (*model.Chat, error) {
	chat1 := model.Chat{
		UserID: userID,
		ShopID: shopID,
	}
	r.DB.Create(&chat1)
	return &chat1, nil
}

func (r *mutationResolver) CreateChat(ctx context.Context, userID int, shopID int, sourceID int, role string, message string, image string, typeArg string) (*model.Chat, error) {
	var chat *model.Chat
	r.DB.Where("user_id = ? AND shop_id = ?", userID, shopID).First(&chat)

	if chat != nil {
		chatDetail := model.ChatDetail{
			ChatID:   chat.ID,
			SourceID: sourceID,
			Role:     role,
			Message:  message,
			Image:    image,
			Type:     typeArg,
		}
		r.DB.Create(&chatDetail)
		return chat, nil
	} else {
		chat1 := model.Chat{
			UserID: userID,
			ShopID: shopID,
		}
		r.DB.Create(&chat1)

		chatDetail := model.ChatDetail{
			ChatID:   chat1.ID,
			SourceID: sourceID,
			Role:     role,
			Message:  message,
			Image:    image,
			Type:     typeArg,
		}
		r.DB.Create(&chatDetail)
		return &chat1, nil
	}
}

func (r *queryResolver) GetChat(ctx context.Context, userID int) ([]*model.Chat, error) {
	var chat []*model.Chat
	r.DB.Where("user_id=?", userID).Find(&chat)
	return chat, nil
}

func (r *queryResolver) GetChatDetail(ctx context.Context, chatID int) ([]*model.ChatDetail, error) {
	var chat []*model.ChatDetail
	r.DB.Where("chat_id=?", chatID).Find(&chat)
	return chat, nil
}
