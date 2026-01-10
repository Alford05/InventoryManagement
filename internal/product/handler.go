package product

import (
	"net/http"
	"strconv"

	"InventoryManagement/pkg/middleware"
	"InventoryManagement/pkg/validation"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)

	group := r.Group("/products")
	{
		group.POST("", middleware.AuthRequired(), middleware.AdminOnly(), func(c *gin.Context) {
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
			var query ListProductsQuery

			if err := c.ShouldBindQuery(&query); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := validation.Validate.Struct(&query); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if query.Page == 0 {
				query.Page = 1
			}
			if query.PageSize == 0 {
				query.PageSize = 20
			}

			products, total, err := service.List(&query)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"data":      products,
				"total":     total,
				"page":      query.Page,
				"page_size": query.PageSize,
			})
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

		group.PUT("/:id", middleware.AuthRequired(), middleware.AdminOnly(), func(c *gin.Context) {
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

		group.DELETE("/:id", middleware.AuthRequired(), middleware.AdminOnly(), func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Param("id"))
			if err := service.Delete(uint(id)); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
				return
			}
			c.Status(http.StatusNoContent)
		})
	}
}
