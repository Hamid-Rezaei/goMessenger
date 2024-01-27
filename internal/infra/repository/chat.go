package repository

import (
	"context"
	"errors"
	"slices"

	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *ChatRepository {
	return &ChatRepository{
		db: db,
	}
}

func (chr *ChatRepository) Create(ctx context.Context, model model.Chat) (*model.Chat, error) {
	tx := chr.db.WithContext(ctx).Begin()

	chat := model
	if err := tx.Create(&chat).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &chat, tx.Commit().Error
}

func (chr *ChatRepository) GetChatList(ctx context.Context, user_id uint) (*[]model.Chat, error) {
	var chats []model.Chat

	if err := chr.db.Find(&chats).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	result := []model.Chat{}
	for _, chat := range chats {
		if slices.Contains(chat.People, user_id) {
			result = append(result, chat)
		}
	}
	return &result, nil
}

func (chr *ChatRepository) GetChat(ctx context.Context, user_id uint, receiver_id uint) (*model.Chat, error) {
	var chats []model.Chat

	if err := chr.db.Find(&chats).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	for _, chat := range chats {
		if slices.Contains(chat.People, user_id) && slices.Contains(chat.People, receiver_id) {
			return &chat, nil
		}
	}
	return nil, nil
}

func (chr *ChatRepository) GetChatById(ctx context.Context, chat_id uint) (*model.Chat, error) {

	var chat = model.Chat{ID: chat_id}

	if err := chr.db.First(&chat).Error; err != nil {
		return nil, err
	}
	return &chat, nil
}

func (chr *ChatRepository) Delete(ctx context.Context, chat_id uint) error {

	var chat = model.Chat{ID: chat_id}
	tx := chr.db.WithContext(ctx).Begin()

	if err := tx.Delete(&chat).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
