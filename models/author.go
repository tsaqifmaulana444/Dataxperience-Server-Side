package models

import "gorm.io/gorm"

type Authors struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	ProfilePhoto string
	Description  string
}
