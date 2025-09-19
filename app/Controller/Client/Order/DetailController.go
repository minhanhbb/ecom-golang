package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

// GET /api/v1/order/detail
func Detail(c *gin.Context) {
	// Lấy user_id từ context (giả sử đã có middleware xác thực và lưu user_id vào context)
	claims, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claimsMap, ok := claims.(map[string]interface{})
	if !ok {
		// Thử ép kiểu sang jwt.MapClaims
		if jwtClaims, ok := claims.(jwt.MapClaims); ok {
			claimsMap = jwtClaims
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims format"})
			return
		}
	}
	userID, ok := claimsMap["user_id"]
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}
	// Tìm order đang pending của user
	var order models.Orders
	if err := database.DB.Where("user_id = ? AND status = 0", userID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	// Lấy tất cả order items của order
	var items []models.OrderItems
	database.DB.Where("order_id = ?", order.ID).Find(&items)

	// Lấy thông tin product cho từng order item
	var resultItems []map[string]interface{}
	for _, item := range items {
		var product models.Product
		database.DB.First(&product, item.ProductID)
		resultItem := map[string]interface{}{
			"ID":        item.ID,
			"OrderId":   item.OrderId,
			"Quantity":  item.Quantity,
			"ProductID": item.ProductID,
			"Price":     item.Price,
			"Product":   product,
		}
		resultItems = append(resultItems, resultItem)
	}
	c.JSON(http.StatusOK, gin.H{
		"order":       order,
		"order_items": resultItems,
	})
}
