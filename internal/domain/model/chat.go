package model

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	ID uint `gorm:"primarykey"`
}

type People struct {
	ChatID uint `gorm:"primarykey:ChatId"`
	UserID uint `gorm:"primarykey:UserId"`
}
