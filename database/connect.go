package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {

	dsn := "postgres://postgres:rootadmin@localhost/os_go?sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

}
