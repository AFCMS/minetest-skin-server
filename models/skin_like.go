package models

import "gorm.io/gorm"

type SkinLike struct {
	gorm.Model
	SkinID    uint
	Skin      Skin `gorm:"not null"`
	AccountID uint
	Account   Account `gorm:"not null"`
}
