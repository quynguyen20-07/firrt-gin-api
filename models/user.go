package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	gorm.Model
}
