package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)

	group := r.Group("/products")
	{
		group.POST("", func(c *gin.Context) {
			var product Product
			if err := c.ShouldBindJSON(&product); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := service.Create(&product); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, product)
		})

		group.GET("", func(c *gin.Context) {
			filters := make(map[string]interface{})
			if category := c.DefaultQuery("category", ""); category != "" {
				filters["category"] = category
			}
			products, _ := service.List(filters)
			c.JSON(http.StatusOK, products)
		})

		group.GET("/:id", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Param("id"))
			product, err := service.Get(uint(id))
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
				return
			}
			c.JSON(http.StatusOK, product)
		})

		group.PUT("/:id", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Param("id"))
			var product Product
			if err := c.ShouldBindJSON(&product); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := service.Update(uint(id), &product); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, product)
		})

		group.DELETE("/:id", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Param("id"))
			if err := service.Delete(uint(id)); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
				return
			}
			c.Status(http.StatusNoContent)
		})
	}
}
