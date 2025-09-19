package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

func Store(c *gin.Context) {
	var input struct {
		Name  string `json:"name" binding:"required"`
		Image string `json:"image" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Categories{Name: input.Name, Image: input.Image}
	result := database.DB.Create(&category)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully", "data": category})
}
