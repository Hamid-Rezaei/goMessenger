package repository

import (
	"context"
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

func (chr *ChatRepository) GetChatList(ctx context.Context, userId uint) (*[]model.Chat, error) {
	var chats []model.Chat
	chr.db.Raw("SELECT c.* FROM chats c INNER JOIN peoples p ON c.id = p.chat_id WHERE p.user_id = ? and c.deleted_at is null", userId).Scan(&chats)

	return &chats, nil
}

func (chr *ChatRepository) GetChat(ctx context.Context, userId uint, receiverId uint) (*model.Chat, error) {
	var chat *model.Chat
	chr.db.Raw("SELECT c.* FROM chats c INNER JOIN (SELECT p1.chat_id as chat_id FROM peoples p1 INNER JOIN peoples p2 ON p1.chat_id = p2.chat_id WHERE p1.user_id = ? and p2.user_id=?) p ON c.id = p.chat_id", userId, receiverId).Scan(&chat)

	return chat, nil
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
