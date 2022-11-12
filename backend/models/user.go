package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID			int       `json:"id"`
	Username	string    `form:"username" json:"username" binding:"required"`
	Email		string    `form:"email" json:"email" binding:"required"`
	Password	string    `form:"password" json:"password" binding:"required"`
	CreatedAt	time.Time `json:"created_at"`
}


func (user *User) HashPassword(password string) {
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    user.Password = string(hashedPassword)
}

func (user *User) ComparePassword(password, hashedPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}