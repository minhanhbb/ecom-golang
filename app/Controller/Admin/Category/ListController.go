package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

func List(c *gin.Context) {
	var categories []models.Categories
	database.DB.Find(&categories)
	c.JSON(http.StatusOK, gin.H{"data": categories})
}
