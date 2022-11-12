package helpers

import (
	"log"
	"strings"

	models "eademir.dev/models"
)

func CheckUserPass(username, password, hashedPassword string) bool {
	user := models.User{}
	err := user.ComparePassword([]byte(password), []byte(hashedPassword))
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}
