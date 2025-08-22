package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Email       string         `json:"email" gorm:"unique;not null"`
	Phone       string         `json:"phone"`
	Address     string         `json:"address"`
	MemberCode  string         `json:"member_code" gorm:"unique;not null"`
	Status      string         `json:"status" gorm:"default:'active'"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Loans       []Loan         `json:"loans,omitempty" gorm:"foreignKey:MemberID"`
}
