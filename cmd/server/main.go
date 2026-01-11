package main

import (
	"fmt"

	"InventoryManagement/configs"
	"InventoryManagement/internal/category"
	"InventoryManagement/internal/order"
	"InventoryManagement/internal/product"
	"InventoryManagement/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.LoadConfig()

	r := gin.Default()

	database := db.Connect(cfg.Database)

	category.RegisterRoutes(r, database)
	product.RegisterRoutes(r, database)
	order.RegisterRoutes(r, database)

	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
