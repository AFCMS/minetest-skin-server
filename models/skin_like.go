package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SkinLike struct {
	gorm.Model
	SkinUUID  uuid.UUID
	Skin      Skin `gorm:"not null"`
	AccountID uint
	Account   Account `gorm:"not null"`
}
