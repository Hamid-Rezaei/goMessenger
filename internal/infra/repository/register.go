package repository

import (
	"context"

	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
)

type UserRepo interface {
	Create(ctx context.Context, model model.User) (uint, error)
	GetByUsernamePhone(_ context.Context, username string, phone string) (*model.User, error)
	GetUserByID(_ context.Context, id uint) (*model.User, error)
	Update(ctx context.Context, user *model.User, id uint) error
	Delete(ctx context.Context, id uint) error
	SearchUser(_ context.Context, keyword string) (*model.User, error)
}

type ContactRepo interface {
	Create(ctx context.Context, model model.Contact) (*model.Contact, error)
	GetList(ctx context.Context, userId uint) (*[]model.Contact, error)
	GetById(_ context.Context, userId uint, contactId uint) (*model.Contact, error)
	Delete(ctx context.Context, userId uint, contactId uint) error
}

type ChatRepo interface {
	Create(ctx context.Context, model model.Chat) (*model.Chat, error)
	GetChatList(ctx context.Context, userId uint) (*[]model.Chat, error)
	GetChatById(ctx context.Context, chatId uint) (*model.Chat, error)
	GetChat(ctx context.Context, userId uint, receiverId uint) (*model.Chat, error)
	Delete(ctx context.Context, chatId uint) error
}

type MessageRepo interface {
	GetMessage(ctx context.Context, chatId uint, messageId uint) (*model.Message, error)
	Delete(ctx context.Context, chatId uint, messageId uint) error
	GetMessagesOfAChat(ctx context.Context, chatId uint) (*[]model.Message, error)
}
