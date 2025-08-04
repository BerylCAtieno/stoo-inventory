package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName    string `gorm:"size:50;not null"`
	LastName     string `gorm:"size:50;not null"`
	Email        string `gorm:"not null"`
	RoleID       uint
	PasswordHash string `gorm:"not null"`
	IsActive     bool   `gorm:"default:true"`
	LastLoginAt  *time.Time
}
