package repository

import (
	"context"

	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (mr *MessageRepository) GetMessage(ctx context.Context, chat_id uint, message_id uint) (*model.Message, error) {

	var message = model.Message{ID: message_id, ChatId: chat_id}

	if err := mr.db.First(&message).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (mr *MessageRepository) Delete(ctx context.Context, chat_id uint, message_id uint) error {

	var message = model.Message{ID: message_id, ChatId: chat_id}
	tx := mr.db.WithContext(ctx).Begin()

	if err := tx.Delete(&message).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (mr *MessageRepository) GetMessagesOfAChat(ctx context.Context, chat_id uint) (*[]model.Message, error) {

	var messages []model.Message

	if err := mr.db.Where("chatid = ?", chat_id).Find(&messages).Error; err != nil {
		return &[]model.Message{}, err
	}
	return &messages, nil
}
