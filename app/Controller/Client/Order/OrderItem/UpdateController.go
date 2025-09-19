package orderitem

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	utils "github.com/minhanhbb/ecom-golang/app/Utils"
	"github.com/minhanhbb/ecom-golang/database"
)

// Update order item quantity and price, recalculate order total, delete order if no items remain
func Update(c *gin.Context) {
	var input struct {
		Quantity int `json:"quantity" binding:"required"`
	}
	id := c.Param("id")
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var orderItem models.OrderItems
	if err := database.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
		return
	}
	// Update quantity
	orderItem.Quantity = input.Quantity
	database.DB.Save(&orderItem)
	// Get order
	var order models.Orders
	if err := database.DB.First(&order, orderItem.OrderId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	// If quantity is 0, delete order item
	if orderItem.Quantity <= 0 {
		database.DB.Delete(&orderItem)
	}
	// Check if any order items remain
	var count int64
	database.DB.Model(&models.OrderItems{}).Where("order_id = ?", orderItem.OrderId).Count(&count)
	if count == 0 {
		database.DB.Delete(&order)
		c.JSON(http.StatusOK, gin.H{"message": "Order deleted (no items left)"})
		return
	}
	// Update order total price
	utils.UpdateOrderTotalPrice(&order)
	c.JSON(http.StatusOK, gin.H{"message": "Order item updated", "data": orderItem})
}
