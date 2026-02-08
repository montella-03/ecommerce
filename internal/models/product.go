package models

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	Stock       int     `gorm:"default:0"`
	CreatedAt   time.Time
}
