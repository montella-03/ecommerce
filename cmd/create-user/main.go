package main

import (
	"errors"
	"flag"
	"log"

	"ecommerce/internal/database"
	"ecommerce/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	email := flag.String("email", "", "email address for the user")
	password := flag.String("password", "", "password for the user")
	flag.Parse()

	if *email == "" || *password == "" {
		log.Fatal("email and password are required")
	}
	if len(*password) < 6 {
		log.Fatal("password must be at least 6 characters")
	}

	database.Connect()
	if err := database.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(*password), 12)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	var user models.User
	err = database.DB.Where("email = ?", *email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = models.User{
			Email:        *email,
			PasswordHash: string(hashed),
		}
		if err := database.DB.Create(&user).Error; err != nil {
			log.Fatal("Failed to create user:", err)
		}
		log.Printf("Created user %s with id %d", user.Email, user.ID)
		return
	}
	if err != nil {
		log.Fatal("Failed to look up user:", err)
	}

	user.PasswordHash = string(hashed)
	if err := database.DB.Save(&user).Error; err != nil {
		log.Fatal("Failed to update user password:", err)
	}
	log.Printf("Updated password for existing user %s with id %d", user.Email, user.ID)
}
