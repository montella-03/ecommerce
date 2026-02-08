package models

import "time"

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	CreatedAt time.Time
	Items     []CartItem `gorm:"foreignKey:CartID"`
}

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint
	ProductID uint
	Quantity  int     `gorm:"check:quantity > 0"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
