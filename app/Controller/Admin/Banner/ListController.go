package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

func List(c *gin.Context) {
	var banners []models.Banners
	database.DB.Find(&banners)
	c.JSON(http.StatusOK, gin.H{"data": banners})
}
