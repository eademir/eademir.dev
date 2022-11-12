package models

type About struct {
	ID       uint8  `json:"id"`
	Username string `json:"username"`
	Detail   string `json:"detail"`
	ImageUrl string `json:"image_url"`
}
