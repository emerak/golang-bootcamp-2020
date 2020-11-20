package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	FirstName string
	LastName  string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
