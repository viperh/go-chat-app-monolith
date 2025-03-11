package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserId  int    `gorm:"not null"`
	Content string `gorm:"not null"`
}
