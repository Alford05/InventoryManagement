package db

import (
	"InventoryManagement/configs"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg configs.DatabaseConfig) *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=inventory port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	fmt.Println("📦 Database connected")
	return db
}
