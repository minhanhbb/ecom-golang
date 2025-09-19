package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

func List(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)

	var result []map[string]interface{}
	for _, p := range products {
		// Lấy các category cho product
		var productCategories []models.ProductCategory
		database.DB.Where("product_id = ?", p.ID).Find(&productCategories)
		var categories []map[string]interface{}
		for _, pc := range productCategories {
			var cat models.Categories
			if err := database.DB.First(&cat, pc.CategoryID).Error; err == nil {
				categories = append(categories, map[string]interface{}{
					"id":   cat.ID,
					"name": cat.Name,
				})
			}
		}
		result = append(result, map[string]interface{}{
			"id":         p.ID,
			"name":       p.Name,
			"desc":       p.Desc,
			"price":      p.Price,
			"images":     p.Images,
			"status":     p.Status,
			"categories": categories,
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}
