package response

import (
	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
)

type ChatWithMessageResponse struct {
	Chat     *model.Chat      `json:"chat"`
	Messages *[]model.Message `json:"messages"`
}

type ChatResponse struct {
	People []uint
	ID     uint `gorm:"primarykey"`
}

type ChatsResponse struct {
	Chats []model.Chat
}

func NewChatResponse(c *model.Chat) *ChatResponse {
	r := new(ChatResponse)
	r.ID = c.ID
	r.People = c.People
	return r
}

func NewChatsResponse(c *[]model.Chat) *ChatsResponse {
	r := new(ChatsResponse)
	r.Chats = *c
	return r
}
