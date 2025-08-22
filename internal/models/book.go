package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Author      string         `json:"author" gorm:"not null"`
	ISBN        string         `json:"isbn" gorm:"unique;not null"`
	Publisher   string         `json:"publisher"`
	Year        int            `json:"year"`
	Category    string         `json:"category"`
	Description string         `json:"description"`
	Stock       int            `json:"stock" gorm:"default:0"`
	Available   int            `json:"available" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Loans       []Loan         `json:"loans,omitempty" gorm:"foreignKey:BookID"`
}
