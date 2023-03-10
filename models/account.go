package models

import (
	"time"
)

type Account struct {
	ID              uint      `json:"id" gorm:"primarykey"`
	Name            string    `json:"name" gorm:"not null;unique"`
	Password        []byte    `json:"-" gorm:"not null"`
	Email           string    `json:"email" gorm:"not null;unique"`
	PermissionLevel int8      `json:"permission_level" gorm:"default:1"`
	Banned          bool      `json:"banned" gorm:"default:false"`
	BanReason       string    `json:"ban_reason" gorm:"default:"`
	CreatedAt       time.Time `json:"created_at"`
	LastConnection  time.Time `json:"last_connection"`
}
