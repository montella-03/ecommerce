package handlers

import (
	"ecommerce/internal/database"
	"ecommerce/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateOrder - Checkout current cart
func CreateOrder(c *gin.Context) {
	userID := c.GetUint("user_id")

	// Begin transaction
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	cart, err := getOrCreateCart(userID)
	if err != nil || len(cart.Items) == 0 {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty or not found"})
		return
	}

	var total float64 = 0
	var orderItems []models.OrderItem

	for _, item := range cart.Items {
		total += item.Product.Price * float64(item.Quantity)

		orderItem := models.OrderItem{
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
			PriceAtTime: item.Product.Price,
		}
		orderItems = append(orderItems, orderItem)

		// Decrease stock
		if err := tx.Model(&models.Product{}).
			Where("id = ?", item.ProductID).
			Update("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Stock update failed"})
			return
		}
	}

	order := models.Order{
		UserID: userID,
		Total:  total,
		Status: "pending",
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Attach items to order
	for i := range orderItems {
		orderItems[i].OrderID = order.ID
		if err := tx.Create(&orderItems[i]).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add order items"})
			return
		}
	}

	// Clear cart
	tx.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{})
	// Optionally delete cart itself or keep empty

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created",
		"order":   order,
	})
}
