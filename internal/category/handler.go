package category

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)

	group := r.Group("/categories")
	{
		group.POST("", func(c *gin.Context) {
			var category Category
			if err := c.ShouldBindJSON(&category); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := service.Create(&category); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, category)
		})

		group.GET("", func(c *gin.Context) {
			categories, _ := service.List()
			c.JSON(http.StatusOK, categories)
		})

		group.GET("/:id", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Param("id"))
			category, err := service.Get(uint(id))
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
				return
			}
			c.JSON(http.StatusOK, category)
		})

		group.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
				return
			}
			if err := service.Delete(uint(id)); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
				return
			}
			c.Status(http.StatusNoContent)
		})
	}
}
