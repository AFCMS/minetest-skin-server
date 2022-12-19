package models

import (
	"gorm.io/gorm"
)

type Skin struct {
	gorm.Model
	UUID        uint   `json:"uuid" gorm:"primaryKey"`
	Description string `json:"description" gorm:"not null"`
	Public      bool   `json:"public" gorm:"not null,default:true"`
	Approved    bool   `json:"approved" gorm:"not null,default:false"`
	OwnerID     uint
	Owner       Account `json:"owner" gorm:"not null"`
	Data        []byte  `json:"data" gorm:"not null"`
}
