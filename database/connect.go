package database

import (
	"onestep/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "postgres://postgres:rootadmin@localhost/os_go?sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = database

	database.AutoMigrate(&models.User{})

}
