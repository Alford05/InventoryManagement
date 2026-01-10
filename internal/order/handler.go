package order

import (
	"net/http"
	"strconv"

	"InventoryManagement/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(db, repo)

	group := r.Group("/orders")
	group.Use(middleware.AuthRequired())
	{
		// POST /orders -> create a new order
		group.POST("", func(c *gin.Context) {
			var req CreateOrderRequest

			// 1️⃣ Bind JSON
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			req.CustomerID = c.GetUint("user_id")

			// 2️⃣ Call service
			order, err := service.CreateOrder(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// 3️⃣ Respond
			c.JSON(http.StatusCreated, order)
		})

		// PUT /order/id/status -> update order status
		group.PUT("/:id/status", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
				return
			}
			var req UpdateOrderStatusRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			order, err := service.UpdateOrderStatus(uint(id), &req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, order)
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
			userID := c.GetUint("user_id")
			role := c.GetString("role")

			if role != "admin" && order.CustomerID != userID {
				c.JSON(http.StatusForbidden, gin.H{"error": "not your order"})
				return
			}

			c.JSON(http.StatusOK, order)
		})

		// Optional: list orders
		group.GET("", func(c *gin.Context) {
			userID := c.GetUint("user_id")
			role := c.GetString("role")

			var customerID uint
			if role != "admin" {
				customerID = userID
			}
			orders, err := service.ListOrders(customerID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, orders)
		})

		group.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
				return
			}

			order, err := service.GetOrder(uint(id))
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
				return
			}
			userID := c.GetUint("user_id")
			role := c.GetString("role")

			if role != "admin" && order.CustomerID != userID {
				c.JSON(http.StatusForbidden, gin.H{"error": "not your order"})
				return
			}
			canceled, err := service.CancelOrder(uint(id))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, canceled)
		})
	}
}
