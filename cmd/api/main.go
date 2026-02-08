package main

import (
	"ecommerce/internal/database"
	"ecommerce/internal/handlers"
	"ecommerce/internal/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	r := gin.Default()

	// Public
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// Protected
	api := r.Group("/api")
	api.Use(middleware.AuthRequired())
	{
		api.GET("/products", handlers.GetProducts)
		api.GET("/products/:id", handlers.GetProduct)
		api.POST("/products", handlers.CreateProduct)
		api.PUT("/products/:id", handlers.UpdateProduct)

		api.GET("/cart", handlers.GetCart)
		api.POST("/cart/items", handlers.AddToCart)
		api.POST("/orders", handlers.CreateOrder)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s", port)
	r.Run(":" + port)
}
