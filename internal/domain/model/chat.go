package model

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	People []uint
}
