package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	var category models.Banners
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	database.DB.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
