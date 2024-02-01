package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name    string `json:"name,omitempty"`
	Chat    Chat   `gorm:"foreignKey:ChatId"`
	Owner   User   `gorm:"foreignKey:ownerId"`
	Members []User `gorm:"foreignKey:UserId" json:"members,omitempty"`
}
