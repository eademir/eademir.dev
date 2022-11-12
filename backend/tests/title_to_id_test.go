package test

import (
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func TestTitleToID(t *testing.T) {
	time := time.Now().UTC()
	title := randStringBytes(150)
	title2 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."

	re, _ := regexp.Compile(`[^a-zA-Z0-9 _]+`)

	blogID := url.QueryEscape(title) + "-" + time.Format("020106")
	blogID2 := strings.ReplaceAll(re.ReplaceAllString(title2, ""), " ", "-")

	assert.Equal(t, 157, utf8.RuneCountInString(blogID))
	assert.Equal(t, "Lorem-ipsum-dolor-sit-amet-consectetur-adipiscing-elit", blogID2)
}

//random string creator as title with fixed length
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func BenchmarkTitleToID(b *testing.B) {
	TestTitleToID(&testing.T{})
}