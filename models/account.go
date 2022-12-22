package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID              int    `json:"id" gorm:"primarykey"`
	Name            string `json:"name" gorm:"not null;unique"`
	Password        []byte `gorm:"not null"`
	Email           string `gorm:"not null;unique"`
	PermissionLevel int8   `gorm:"default:1"`
	CreationDate    int64  `gorm:"not null"`
	LastConnection  int64
}
