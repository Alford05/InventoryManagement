package testutils

import (
	"InventoryManagement/internal/order"
	"InventoryManagement/internal/product"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&product.Product{},
		&order.Order{},
		&order.OrderItem{},
	)
	return db
}
