package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Title      string
	Content    string
	CategoryID uint
	Category   Categories `gorm:"foreignKey:CategoryID"`
	AuthorID   uint
	Author     Authors    `gorm:"foreignKey:AuthorID"`
}
