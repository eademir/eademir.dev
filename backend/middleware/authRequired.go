package middleware

import (
	"eademir.dev/globals"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	if user := session.Get(globals.Userkey); user == nil {
		log.Println("User not logged in")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
}
