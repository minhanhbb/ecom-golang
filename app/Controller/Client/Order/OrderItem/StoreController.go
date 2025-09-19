package orderitem

import (
	"fmt"
	"github.com/gin-gonic/gin"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
	"net/http"
)

func Store(c *gin.Context) {
	var input struct {
		UserID    string `json:"user_id" binding:"required"`
		ProductID int    `json:"product_id" binding:"required"`
		Quantity  int    `json:"quantity" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Tìm order đang pending của user
	var order models.Orders
	if err := database.DB.Where("user_id = ? AND status = 0", input.UserID).First(&order).Error; err != nil {
		// Nếu chưa có order thì tạo mới
		order = models.Orders{
			UserID:     input.UserID,
			TotalPrice: "0",
			Status:     0,
		}
		database.DB.Create(&order)
	}
	// Kiểm tra order_item đã tồn tại chưa
	var orderItem models.OrderItems
	if err := database.DB.Where("order_id = ? AND product_id = ?", int(order.ID), input.ProductID).First(&orderItem).Error; err == nil {
		// Nếu đã có thì tăng số lượng
		orderItem.Quantity = orderItem.Quantity + input.Quantity
		database.DB.Save(&orderItem)
		// Cập nhật lại tổng tiền
		updateOrderTotalPrice(&order)
		c.JSON(http.StatusOK, gin.H{"message": "Order item updated", "data": orderItem})
		return
	}
	// Nếu chưa có thì tạo mới
	var product models.Product
	if err := database.DB.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	orderItem = models.OrderItems{
		OrderId:   int(order.ID),
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Price:     product.Price,
	}
	database.DB.Create(&orderItem)
	// Cập nhật lại tổng tiền
	updateOrderTotalPrice(&order)
	c.JSON(http.StatusOK, gin.H{"message": "Order item created", "data": orderItem})

}

// Hàm cập nhật tổng tiền cho order
func updateOrderTotalPrice(order *models.Orders) {
	var items []models.OrderItems
	database.DB.Where("order_id = ?", int(order.ID)).Find(&items)
	var total float64 = 0
	for _, item := range items {
		total += float64(item.Quantity) * item.Price
	}
	order.TotalPrice = fmt.Sprintf("%.2f", total)
	database.DB.Save(order)
}
