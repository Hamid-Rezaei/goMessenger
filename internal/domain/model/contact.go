package model

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	UserId         uint   `gorm:"primaryKey;autoIncrement:false"`
	User           User   `gorm:"foreignKey:UserId"`
	ContactId      uint   `gorm:"primaryKey;autoIncrement:false"`
	Contact        User   `gorm:"foreignKey:ContactId"`
	ContactName    string `json:"contact_name,omitempty"`
	ShowNumber     bool   `json:"show_number,omitempty"`
	ShowProfilePic bool   `json:"show_profile_pic,omitempty"`
}
