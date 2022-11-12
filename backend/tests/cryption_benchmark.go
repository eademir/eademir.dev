package test

import (
	"crypto/rand"
	"encoding/base32"
	"testing"

	models "eademir.dev/models"
)

func Benchmark20DigitsEncryption(b *testing.B) {
	user := models.User{}
	password := generatePassword(20)

	user.HashPassword(password)

}

func generatePassword(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}
