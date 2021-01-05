package model

import (
	"gorm.io/gorm"
)

type Platform struct {
	gorm.Model
	ID 			uint	`gorm:"id"`
	Name 		string	`gorm:"name"`
}

func (platform *Platform) TableName() string {
	return "platform"
}