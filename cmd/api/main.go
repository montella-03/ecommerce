package main

import (
	"ecommerce/internal/database"
	"ecommerce/internal/handlers"
	"ecommerce/internal/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	r.Static("/assets", "./web/assets")
	r.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	})

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
		api.DELETE("/products/:id", handlers.DeleteProduct)
		api.GET("/users", handlers.GetUsers)

		api.GET("/cart", handlers.GetCart)
		api.POST("/cart/items", handlers.AddToCart)
		api.POST("/orders", handlers.CreateOrder)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s", port)
	err := r.Run(":" + port)
	if err != nil {
		return
	}
}
