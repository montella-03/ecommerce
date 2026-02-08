package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint
	Total  float64 `gorm:"type:decimal(10,2)"`
	Status string  `gorm:"default:'pending'"`
	Items  []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID     uint
	ProductID   uint
	Quantity    int
	PriceAtTime float64 `gorm:"type:decimal(10,2)"`
}
