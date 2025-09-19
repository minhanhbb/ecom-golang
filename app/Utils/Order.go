package utils

import (
	"fmt"

	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

// Update order's total price based on its items
func UpdateOrderTotalPrice(order *models.Orders) {
	var items []models.OrderItems
	database.DB.Where("order_id = ?", int(order.ID)).Find(&items)
	var total float64 = 0
	for _, item := range items {
		total += float64(item.Quantity) * item.Price
	}
	order.TotalPrice = fmt.Sprintf("%.2f", total)
	database.DB.Save(order)
}
