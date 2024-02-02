package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name    string `json:"name,omitempty"`
	ChatID  uint   // Foreign key for the Chat model
	Chat    Chat   `gorm:"foreignKey:ChatID"`
	OwnerID uint   // Foreign key for the User model
	Owner   *User  `gorm:"foreignKey:OwnerID"`
	Members []User `gorm:"many2many:group_members;" json:"members,omitempty"`
}
