package models

import "gorm.io/gorm"

type Product struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `gorm:"unique;not null"`
	Desc  string  `gorm:"not null"`
	Image *string `gorm:"NULL"`
	gorm.Model
}
