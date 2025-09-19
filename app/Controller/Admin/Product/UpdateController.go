package product

import (
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

func Update(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Name        string   `json:"name"`
		Desc        string   `json:"desc"`
		Price       float64  `json:"price"`
		CategoryIDs []uint   `json:"category_id"`
		Images      []string `json:"images"`
		Status      int      `json:"status"`
	}
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Name != "" {
		product.Name = input.Name
	}
	if input.Desc != "" {
		product.Desc = input.Desc
	}
	if input.Price != 0 {
		product.Price = input.Price
	}
	// Cập nhật images
	if input.Images != nil {
		imagesJson, err := json.Marshal(input.Images)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid images format"})
			return
		}
		product.Images = string(imagesJson)
	}
	// Cập nhật lại các category_id
	if input.CategoryIDs != nil {
		// Xóa các liên kết cũ
		database.DB.Where("product_id = ?", product.ID).Delete(&models.ProductCategory{})
		// Thêm liên kết mới
		for _, catID := range input.CategoryIDs {
			pc := models.ProductCategory{ProductID: product.ID, CategoryID: catID}
			database.DB.Create(&pc)
		}
	}
	if input.Status != 0 {
		product.Status = input.Status
	}
	database.DB.Save(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product updated", "data": product})
}
