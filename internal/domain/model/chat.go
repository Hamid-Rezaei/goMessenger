package model

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	People    []uint
	CreatedAt time.Time `json:"created_at,omitempty"`
}
