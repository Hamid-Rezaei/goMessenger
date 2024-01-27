package model

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey"`
	People    []uint64
	CreatedAt time.Time `json:"created_at,omitempty"`
}
