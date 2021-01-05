package model

import (
	"gorm.io/gorm"
)

type Games struct {
	gorm.Model
	ID 			uint	`gorm:"id"`
	Name 		string	`gorm:"name"`
	Description string	`gorm:"description"`
	IDPlatform 	uint	`gorm:"id_platform"`
	Plus		uint	`gorm:"plus"`
}

func (games *Games) TableName() string {
	return "games"
}