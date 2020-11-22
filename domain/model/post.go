package model

import "time"

type Post struct {
	ID          uint
	UserID      uint
	User        User
	Title       string
	Body        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"default:null"`
	PublishedAt time.Time `gorm:"default:null"`
}
