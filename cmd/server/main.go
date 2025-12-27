package main

import (
	"inventory-api/internal/category"
	"inventory-api/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database := db.Connect()

	category.RegisterRoutes(r, database)

	r.Run(":8080")
}
