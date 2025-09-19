package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

// DELETE /api/v1/order/delete
func Delete(c *gin.Context) {
	claims, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims format"})
		return
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
	// Xóa toàn bộ orderitem
	database.DB.Where("order_id = ?", order.ID).Delete(&models.OrderItems{})
	// Xóa order
	database.DB.Delete(&order)
	c.JSON(http.StatusOK, gin.H{"message": "Order and order items deleted"})
}
