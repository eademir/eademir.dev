package controllers

import (
	"eademir.dev/globals"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	store := cookie.NewStore(globals.Secret)
	r.Use(sessions.Sessions(globals.Userkey, store))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return r
}
