package model

import (
	"time"
)

type Users []User

type User struct {
	ID uint	`gorm:"primary_key"`
	FirstName string `gorm:"not null"`
	LastName string `gorm:"not null" validate:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
