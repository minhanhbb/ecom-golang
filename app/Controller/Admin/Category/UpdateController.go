package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

func Update(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Name  string `json:"name"`
		Image string `json:"image"`
	}
	var category models.Categories
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Name != "" {
		category.Name = input.Name
	}
	if input.Image != "" {
		category.Image = input.Image
	}
	database.DB.Save(&category)
	c.JSON(http.StatusOK, gin.H{"message": "Category updated", "data": category})
}
