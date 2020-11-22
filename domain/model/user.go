package model

import (
	"time"
)

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"default:null"`
}
