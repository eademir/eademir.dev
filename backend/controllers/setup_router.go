package controllers

import (
	"eademir.dev/globals"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://eademir.dev", "https://admin.eademir.dev"}
	config.AllowMethods = []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"}

	r.Use(cors.New(config))

	store := cookie.NewStore(globals.Secret)
	r.Use(sessions.Sessions(globals.Userkey, store))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return r
}
