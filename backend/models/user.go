package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	NewPassword string    `json:"new_password"`
	CreatedAt   time.Time `json:"created_at"`
}

func (user *User) HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func (user *User) ComparePassword(password, hashedPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
