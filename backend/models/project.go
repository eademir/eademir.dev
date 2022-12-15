package models

import "time"

type Project struct {
	ID          uint8     `json:"id" db:"id"`
	Username    string    `json:"username" db:"username"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Images      []string  `json:"images" db:"images"`
	Links       []string  `json:"links" db:"links"`
	Date        time.Time `json:"date" db:"date"`
}
