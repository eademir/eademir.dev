package models

type Project struct {
	ID          uint8  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Link        string `json:"link"`
}
