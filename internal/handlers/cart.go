package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/montella-03/ecommerce/internal/database"
	"github.com/montella-03/ecommerce/internal/models"
)

// getOrCreateCart - helper (not exported)
func getOrCreateCart(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := database.DB.Where("user_id = ?", userID).Preload("Items.Product").First(&cart).Error
	if err == nil {
		return &cart, nil
	}

	// Create new cart
	cart = models.Cart{UserID: userID}
	if err := database.DB.Create(&cart).Error; err != nil {
		return nil, err
	}

	// Reload with items (empty)
	database.DB.Preload("Items.Product").First(&cart, cart.ID)
	return &cart, nil
}

// GetCart
func GetCart(c *gin.Context) {
	userID := c.GetUint("user_id")

	cart, err := getOrCreateCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

// AddToCart
func AddToCart(c *gin.Context) {
	userID := c.GetUint("user_id")

	type AddInput struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required,gt=0"`
	}

	var input AddInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check product exists and has stock
	var product models.Product
	if err := database.DB.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if product.Stock < input.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock"})
		return
	}

	cart, err := getOrCreateCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cart error"})
		return
	}

	// Check if already in cart → update quantity
	var existingItem models.CartItem
	err = database.DB.Where("cart_id = ? AND product_id = ?", cart.ID, input.ProductID).First(&existingItem).Error
	if err == nil {
		// Update quantity
		newQty := existingItem.Quantity + input.Quantity
		if product.Stock < newQty {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock"})
			return
		}

		database.DB.Model(&existingItem).Update("quantity", newQty)
	} else {
		// New item
		item := models.CartItem{
			CartID:    cart.ID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		database.DB.Create(&item)
	}

	// Reload cart with items
	database.DB.Preload("Items.Product").First(cart, cart.ID)

	c.JSON(http.StatusOK, gin.H{"cart": cart})
}
