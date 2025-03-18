package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	User []User `gorm:"many2many:room_users;"`
}
