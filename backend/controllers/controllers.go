package controllers

import (
	"database/sql"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"net/http"

	"eademir.dev/globals"
	"eademir.dev/helpers"
	models "eademir.dev/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Create a blog
func CreateBlog(c *gin.Context, db *sql.DB) {
	blog := models.Blog{}

	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	c.BindJSON(&blog)

	//title to id
	time := time.Now()
	re, _ := regexp.Compile(`[^a-zA-Z0-9 _]+`)
	blog.ID = strings.ReplaceAll(re.ReplaceAllString(strings.Trim(blog.Title, " "), ""), " ", "-") + "-" + time.Format("020106")

	sqlStatement := `INSERT INTO blogs (id, title, description, content, keywords, username) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err := db.QueryRow(sqlStatement, blog.ID, blog.Title, blog.Description, blog.Content, blog.Keywords, user).Scan(&blog.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
}

// Update a blog
func UpdateBlog(c *gin.Context, db *sql.DB) {
	blog := models.Blog{}

	//get id from URL
	id := c.Param("id")

	c.BindJSON(&blog)

	blog.UpdatedAt = time.Now()

	sqlStatement := `Update blogs SET  title = $2, description = $3, content = $4, keywords = $5, edited_at = $6 WHERE id = $1;`

	_, err := db.Exec(sqlStatement, id, blog.Title, blog.Description, blog.Content, blog.Keywords, blog.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
}

// Delete handler by id from database
func DeleteHandler(c *gin.Context, db *sql.DB, table string) {

	//get id from URL and convert to int
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		return
	}

	sqlStatement := `Delete FROM ` + table + ` WHERE id = $1;`

	//execute sql statement
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
}

// get a blog from database
func ReadBlog(c *gin.Context, db *sql.DB) {
	blog := models.Blog{}

	//get id from URL and convert to int
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		return
	}
	sqlStatement := `SELECT * FROM blogs WHERE id = $1;`

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&blog.ID, &blog.Title, &blog.Description, &blog.Content, &blog.Keywords, &blog.CreatedAt, &blog.UpdatedAt, &blog.IsShared)

	// if no rows are returned, then the blog does not exist
	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{
			"message": "No rows were returned!",
		})
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"blog": blog,
		})
	default:
		panic(err)
	}
}

// Get all blogs from database function
func ReadBlogs(c *gin.Context, db *sql.DB) {
	blogs := []models.Blog{}

	sqlStatement := `SELECT * FROM blogs;`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	defer rows.Close()

	//iterate over rows
	for rows.Next() {
		blog := models.Blog{}

		//scan rows into blog struct
		err = rows.Scan(&blog.ID, &blog.Username, &blog.Title, &blog.Description, &blog.Content, &blog.Keywords, &blog.CreatedAt, &blog.UpdatedAt, &blog.IsShared, &blog.ImageURL)
		if err != nil {
			panic(err)
		}

		//append blog to blogs if it is shared
		if blog.IsShared {
			blogs = append(blogs, blog)
		}
	}

	//return blogs json with status code 200
	c.JSON(http.StatusOK, gin.H{
		"blogs": blogs,
	})
}

// get about from database
func ReadAbout(c *gin.Context, db *sql.DB) {
	about := models.About{}

	sqlStatement := `SELECT * FROM about;`

	row := db.QueryRow(sqlStatement)
	err := row.Scan(&about.ID, &about.Username, &about.ImageUrl, &about.Detail)

	// if no rows are returned, then the about does not exist
	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{
			"message": "No rows were returned!",
		})
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"about": about,
		})
	default:
		panic(err)
	}
}

// Update about in database
func UpdateAbout(c *gin.Context, db *sql.DB) {
	about := models.About{}
	c.BindJSON(&about)

	sqlStatement := `Update about SET detail = $2, image_url = $3 WHERE id = $1;`

	//execute sql statement
	_, err := db.Exec(sqlStatement, about.ID, about.Detail, about.ImageUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
}

// Read all projetcs from database
func ReadProjects(c *gin.Context, db *sql.DB) {
	projects := []models.Project{}

	sqlStatement := `SELECT * FROM projects;`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	defer rows.Close()

	//iterate over rows
	for rows.Next() {
		project := models.Project{}

		//scan rows into project struct
		err = rows.Scan(&project.ID, &project.Title, &project.Description, &project.ImageUrl, &project.Link)
		if err != nil {
			panic(err)
		}

		//append project to projects
		projects = append(projects, project)
	}

	//return projects json with status code 200
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

// Create a project
func CreateProject(c *gin.Context, db *sql.DB) {
	project := models.Project{}
	c.BindJSON(&project)

	sqlStatement := `INSERT INTO projects (title, description, image_url, link) VALUES ($1, $2, $3, $4) RETURNING id`

	err := db.QueryRow(sqlStatement, project.Title, project.Description, project.ImageUrl, project.Link).Scan(&project.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
}

// Update a project
func UpdateProject(c *gin.Context, db *sql.DB) {
	project := models.Project{}
	c.BindJSON(&project)

	sqlStatement := `Update projects SET title = $2, description = $3, image_url = $4, link = $5 WHERE id = $1;`

	//execute sql statement
	_, err := db.Exec(sqlStatement, project.ID, project.Title, project.Description, project.ImageUrl, project.Link)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
}

// Read a project
func ReadProject(c *gin.Context, db *sql.DB) {
	project := models.Project{}

	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		return
	}

	sqlStatement := `SELECT * FROM projects WHERE id = $1;`

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&project.ID, &project.Title, &project.Description, &project.ImageUrl, &project.Link)

	// if no rows are returned, then the project does not exist
	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{
			"message": "No rows were returned!",
		})
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"project": project,
		})
	default:
		panic(err)
	}
}

// create user
func CreateUser(c *gin.Context, db *sql.DB) {
	user := models.User{}
	c.BindJSON(&user)

	sqlStatement := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`

	user.HashPassword(user.Password)

	err := db.QueryRow(sqlStatement, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
}

// user login post
func LoginPost(c *gin.Context, db *sql.DB) {
	user := models.User{}
	session := sessions.Default(c)
	c.ShouldBindJSON(&user)

	var hashedPassword string

	sqlStatement := `SELECT password FROM users WHERE username = $1;`

	row := db.QueryRow(sqlStatement, user.Username)

	if err := row.Scan(&hashedPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User is not found"})
		return
	}

	if helpers.EmptyUserPass(user.Username, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Parameters can't be empty"})
		return
	}

	if !helpers.CheckUserPass(user.Username, user.Password, hashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect username or password"})
		return
	}

	session.Set(globals.Userkey, user.Username)
	session.Save()

	c.Redirect(http.StatusMovedPermanently, "/admin")
}

// user login
func Login(c *gin.Context, db *sql.DB) {
	user := models.User{}
	c.ShouldBindJSON(&user)

	session := sessions.Default(c)
	username := session.Get(globals.Userkey)
	sqlStatement := `SELECT * FROM users WHERE username = $1;`

	row := db.QueryRow(sqlStatement, username)
	_ = row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.Password)

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"created_at": user.CreatedAt,
	})
}

// logout
func Logout(c *gin.Context, db *sql.DB) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	log.Println("logging out user:", user)
	if user == nil {
		log.Println("Invalid session token")
		return
	}

	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	session.Save()

	if err := session.Get(globals.Userkey); err != nil {
		log.Println("Failed to delete the session")
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}
