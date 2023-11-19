package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(postgres.Open("host=localhost user=postgres dbname=dataxperience_server_2023111909213 password=Saturnuss1; sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// models auto migrate
	database.AutoMigrate(&Categories{}, &News{}, &Authors{}, &Clients{})

	DB = database
}
