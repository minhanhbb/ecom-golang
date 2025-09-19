package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

func Detail(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	// Lấy các category cho product
	var productCategories []models.ProductCategory
	database.DB.Where("product_id = ?", product.ID).Find(&productCategories)
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
	c.JSON(http.StatusOK, gin.H{"data": map[string]interface{}{
		"id":         product.ID,
		"name":       product.Name,
		"desc":       product.Desc,
		"price":      product.Price,
		"images":     product.Images,
		"status":     product.Status,
		"categories": categories,
	}})
}
