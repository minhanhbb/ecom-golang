package product

import (
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

func Store(c *gin.Context) {
	var input struct {
		Name        string   `json:"name" binding:"required"`
		Desc        string   `json:"desc" binding:"required"`
		Price       float64  `json:"price" binding:"required"`
		CategoryIDs []uint   `json:"category_id" binding:"required"`
		Images      []string `json:"images"`
		Status      int      `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	imagesJson, err := json.Marshal(input.Images)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid images format"})
		return
	}
	product := models.Product{
		Name:   input.Name,
		Desc:   input.Desc,
		Price:  input.Price,
		Images: string(imagesJson),
		Status: input.Status,
	}
	result := database.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	// Lưu các category_id vào bảng product_categories
	for _, catID := range input.CategoryIDs {
		pc := models.ProductCategory{ProductID: product.ID, CategoryID: catID}
		database.DB.Create(&pc)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "data": product})
}
