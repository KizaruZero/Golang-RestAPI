package models

import (
	"time"

	"gorm.io/gorm"
)

// Book struct
type Book struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string         `gorm:"type:varchar(255);not null" json:"title"`
	Author    string         `gorm:"type:varchar(255);not null" json:"author"`
	Price     int            `gorm:"not null" json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete field
}
