package routes

import (
	"github.com/gin-gonic/gin"

	"database/sql"

	controllers "eademir.dev/controllers"
)

func PublicRoutes(g *gin.RouterGroup, db *sql.DB) {

	//create user
	//g.POST("/create", func(c *gin.Context) {
	//	controllers.CreateUser(c, db)
	//})

	// get all blogs
	g.GET("/blogs", func(c *gin.Context) {
		controllers.ReadBlogs(c, db, false)
	})

	// get a blog
	g.GET("/blogs/:id", func(c *gin.Context) {
		controllers.ReadBlog(c, db)
	})

	// get about me
	g.GET("/", func(c *gin.Context) {
		controllers.ReadAbout(c, db)
	})

	// get all projects
	g.GET("/projects", func(c *gin.Context) {
		controllers.ReadProjects(c, db)
	})

	// admin login
	g.POST("/login", func(c *gin.Context) {
		controllers.LoginPost(c, db)
	})

}

func PrivateRoutes(g *gin.RouterGroup, db *sql.DB) {

	// admin get all blogs
	g.GET("/blogs", func(c *gin.Context) {
		controllers.ReadBlogs(c, db, true)
	})

	// admin post a blog
	g.POST("/blogs", func(c *gin.Context) {
		controllers.CreateBlog(c, db)
	})

	// admin get a blog
	g.GET("/blogs/:id", func(c *gin.Context) {
		controllers.ReadBlog(c, db)
	})

	// admin update a blog
	g.PUT("/blogs/:id", func(c *gin.Context) {
		controllers.UpdateBlog(c, db)
	})

	// admin delete a blog
	g.DELETE("/blogs/:id", func(c *gin.Context) {
		controllers.DeleteHandler(c, db, "blogs")
	})

	//admin get all projects
	g.GET("/projects", func(c *gin.Context) {
		controllers.ReadProjects(c, db)
	})

	// admin post a project
	g.POST("/projects", func(c *gin.Context) {
		controllers.CreateProject(c, db)
	})

	// admin get a project
	g.GET("/projects/:id", func(c *gin.Context) {
		controllers.ReadProject(c, db)
	})

	// admin update a project
	g.PUT("/projects/:id", func(c *gin.Context) {
		controllers.UpdateProject(c, db)
	})

	// admin delete a project
	g.DELETE("/projects/:id", func(c *gin.Context) {
		controllers.DeleteHandler(c, db, "projects")
	})

	// admin get about me
	g.GET("/about", func(c *gin.Context) {
		controllers.ReadAbout(c, db)
	})

	// admin update about me
	g.PUT("/about", func(c *gin.Context) {
		controllers.UpdateAbout(c, db)
	})

	// admin logout
	g.POST("/logout", func(c *gin.Context) {
		controllers.Logout(c, db)
	})

	g.GET("/dashboard", func(c *gin.Context) {
		controllers.Dashboard(c, db)
	})

}
