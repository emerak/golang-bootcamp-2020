package model

import "time"

type Post struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	PublishedAt time.Time `json:"deleted_at"`
}
