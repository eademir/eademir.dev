package models

import (
	"time"
)

type Blog struct {
	ID          string    `json:"id" db:"id"`
	Username    string    `json:"username" db:"username"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Content     string    `json:"content" db:"content"`
	Keywords    string    `json:"keywords" db:"keywords"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   any       `json:"updated_at" db:"updated_at"`
	IsShared    bool      `json:"is_shared" db:"is_shared"`
	ImageURL    string    `json:"image_url" db:"image_url"`
}
