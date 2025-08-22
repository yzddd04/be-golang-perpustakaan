package models

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	BookID       uint           `json:"book_id" gorm:"not null"`
	MemberID     uint           `json:"member_id" gorm:"not null"`
	LoanDate     time.Time      `json:"loan_date" gorm:"not null"`
	DueDate      time.Time      `json:"due_date" gorm:"not null"`
	ReturnDate   *time.Time     `json:"return_date"`
	Status       string         `json:"status" gorm:"default:'borrowed'"`
	Fine         float64        `json:"fine" gorm:"default:0"`
	Notes        string         `json:"notes"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
	Book         Book           `json:"book,omitempty" gorm:"foreignKey:BookID"`
	Member       Member         `json:"member,omitempty" gorm:"foreignKey:MemberID"`
}
