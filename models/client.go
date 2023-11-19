package models

import "gorm.io/gorm"

type Clients struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	ProfilePhoto string
	Country      string
}
