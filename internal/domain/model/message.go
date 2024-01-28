package model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID         uint   `gorm:"primarykey"`
	ChatId     uint   `json:"chat_id,omitempty"`
	Chat       Chat   `gorm:"foreignKey:ChatId"`
	SenderId   uint   `json:"sender_id,omitempty"`
	Sender     User   `gorm:"foreignKey:SenderId"`
	ReceiverId uint   `json:"receiver_id,omitempty"`
	Receiver   User   `gorm:"foreignKey:ReceiverId"`
	Content    string `json:"content,omitempty"`
}
