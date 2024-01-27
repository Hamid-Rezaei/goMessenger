package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID         uint64    `gorm:"primaryKey"`
	ChatId     uint64    `json:"chatid,omitempty"`
	Chat       Chat      `gorm:"foreignKey:ChatId"`
	SenderId   uint64    `json:"senderid,omitempty"`
	Sender     User      `gorm:"foreignKey:SenderId"`
	ReceiverId uint64    `json:"receiverid,omitempty"`
	Receiver   User      `gorm:"foreignKey:ReceiverId"`
	Content    string    `json:"content,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
