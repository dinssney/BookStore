package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=localhost user=user password=password dbname=bookdb port=5432 sslmode=disable")
	var err error
	Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}
