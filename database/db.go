package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDataBase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("Database connection error:", err)
	}
}
