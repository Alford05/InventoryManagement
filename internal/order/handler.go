package order

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(db, repo)

	group := r.Group("/orders")
	{
		// POST /orders -> create a new order
		group.POST("", func(c *gin.Context) {
			var req CreateOrderRequest

			// 1️⃣ Bind JSON
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// 2️⃣ Call service
			order, err := service.CreateOrder(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// 3️⃣ Respond
			c.JSON(http.StatusCreated, order)
		})

		// GET /orders/:id -> view single order
		group.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
				return
			}

			order, err := service.GetOrder(uint(id))
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
				return
			}

			c.JSON(http.StatusOK, order)
		})

		// Optional: list orders
		group.GET("", func(c *gin.Context) {
			customerIDStr := c.Query("customer_id")
			var customerID uint
			if customerIDStr != "" {
				id, err := strconv.Atoi(customerIDStr)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer_id"})
					return
				}
				customerID = uint(id)
			}
			orders, err := service.ListOrders(customerID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, orders)
		})
	}
}
