package model

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	UserId         uint64 `gorm:"primaryKey;autoIncrement:false"`
	User           User   `gorm:"foreignKey:UserId"`
	ContactId      uint64 `gorm:"primaryKey;autoIncrement:false"`
	Contact        User   `gorm:"foreignKey:ContactId"`
	ContactName    string `json:"contactname,omitempty"`
	ShowNumber     bool   `json:"shownumber,omitempty"`
	ShowProfilePic bool   `json:"showprofilepic,omitempty"`
}
