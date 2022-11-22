package models

type BlogGlance struct {
	ID          string `json:"id" db:"id"`
	ImageURL    string `json:"image_url" db:"image_url"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	IsShared    bool   `json:"is_shared" db:"is_shared"`
}
