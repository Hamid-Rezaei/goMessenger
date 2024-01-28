package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey"`
	ChatId     uint      `json:"chatid,omitempty"`
	Chat       Chat      `gorm:"foreignKey:ChatId"`
	SenderId   uint      `json:"senderid,omitempty"`
	Sender     User      `gorm:"foreignKey:SenderId"`
	ReceiverId uint      `json:"receiverid,omitempty"`
	Receiver   User      `gorm:"foreignKey:ReceiverId"`
	Content    string    `json:"content,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
