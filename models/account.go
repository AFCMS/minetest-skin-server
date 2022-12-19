package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Id              int    `json:"id" gorm:"primaryKey"`
	Name            string `json:"name" gorm:"not null"`
	Password        string `gorm:"not null"`
	Email           string `gorm:"not null;unique"`
	PermissionLevel int8   `gorm:"default:1"`
	CreationDate    int64  `gorm:"not null"`
	LastConnection  int64
}
