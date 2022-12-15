package controllers

import (
	"database/sql"
	"fmt"
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
	"github.com/lib/pq"
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

	sqlStatement := `INSERT INTO blogs (id, title, description, content, keywords, username, is_shared, image_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := db.QueryRow(sqlStatement, blog.ID, blog.Title, blog.Description, blog.Content, blog.Keywords, user, blog.IsShared, blog.ImageURL).Scan(&blog.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// Update a blog
func UpdateBlog(c *gin.Context, db *sql.DB) {
	blog := models.Blog{}

	//get id from URL
	id := c.Param("id")

	c.ShouldBindJSON(&blog)

	time := time.Now()
	re, _ := regexp.Compile(`[^a-zA-Z0-9 _]+`)
	blog.ID = strings.ReplaceAll(re.ReplaceAllString(strings.Trim(blog.Title, " "), ""), " ", "-") + "-" + time.Format("020106")

	sqlStatement := `UPDATE blogs SET id = $7,  title = $2, description = $3, content = $4, keywords = $5, edited_at = $6, is_shared = $8, image_url = $9 WHERE id = $1;`

	_, err := db.Exec(sqlStatement, id, blog.Title, blog.Description, blog.Content, blog.Keywords, time, blog.ID, blog.IsShared, blog.ImageURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// Delete handler by id from database
func DeleteHandler(c *gin.Context, db *sql.DB, table string) {

	//get id from URL and convert to int
	id := c.Param("id")

	sqlStatement := `Delete FROM ` + table + ` WHERE id = $1;`

	//execute sql statement
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// get a blog from database
func ReadBlog(c *gin.Context, db *sql.DB) {
	blog := models.Blog{}

	id := c.Param("id")

	sqlStatement := `SELECT * FROM blogs WHERE id = $1;`

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&blog.ID, &blog.Username, &blog.Title, &blog.Description, &blog.Content, &blog.Keywords, &blog.CreatedAt, &blog.UpdatedAt, &blog.IsShared, &blog.ImageURL)

	// if no rows are returned, then the blog does not exist
	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{
			"message": "No rows were returned!",
		})
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"id":          blog.ID,
			"imageURL":    blog.ImageURL,
			"title":       blog.Title,
			"description": blog.Description,
			"content":     blog.Content,
			"keywords":    blog.Keywords,
			"createdAt":   blog.CreatedAt,
			"updatedAt":   blog.UpdatedAt,
		})
	default:
		panic(err)
	}
}

// Get all blogs from database function
func ReadBlogs(c *gin.Context, db *sql.DB, isAdmin bool, offset, limit int) {
	blogs := []models.BlogGlance{}
	var count int

	sqlStatement := `SELECT id, title, description, is_shared, image_url FROM blogs ORDER BY id LIMIT $1 OFFSET $2;`

	if err := db.QueryRow("SELECT COUNT(*) FROM blogs").Scan(&count); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(sqlStatement, limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	//iterate over rows
	for rows.Next() {
		blogGlance := models.BlogGlance{}

		//scan rows into blog struct
		err = rows.Scan(
			&blogGlance.ID,
			&blogGlance.Title,
			&blogGlance.Description,
			&blogGlance.IsShared,
			&blogGlance.ImageURL,
		)

		switch err {
		case sql.ErrNoRows:
			c.JSON(http.StatusOK, gin.H{
				"message": "No rows were returned!",
			})
		case nil:
			//append blog to blogs if it is shared or admin logged in
			if blogGlance.IsShared || isAdmin {
				blogs = append(blogs, blogGlance)
			}
		default:
			panic(err)
		}
	}

	//return blogs json with status code 200
	c.JSON(http.StatusOK, gin.H{
		"blogs":      blogs,
		"totalBlogs": count,
	})

	defer rows.Close()
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
	_, err := db.Exec(sqlStatement, 1, about.Detail, about.ImageUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
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
	err := row.Scan(&project.ID, &project.Username, &project.Title, &project.Description, &project.Date, pq.Array(&project.Images), pq.Array(&project.Links))

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
		err = rows.Scan(&project.ID, &project.Username, &project.Title, &project.Description, &project.Date, pq.Array(&project.Images), pq.Array(&project.Links))
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

	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	sqlStatement := `INSERT INTO projects (username, title, description, date, images, links) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err := db.QueryRow(sqlStatement, user, project.Title, project.Description, project.Date, pq.Array(project.Images), pq.Array(project.Links)).Scan(&project.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// Update a project
func UpdateProject(c *gin.Context, db *sql.DB) {
	project := models.Project{}

	fmt.Sscan(c.Param("id"), &project.ID)

	c.ShouldBindJSON(&project)

	sqlStatement := `UPDATE projects SET username = $2, title = $3, date = $4, description = $5, images = $6, links = $7 WHERE id = $1;`

	//execute sql statement
	_, err := db.Exec(sqlStatement, project.ID, project.Username, project.Title, project.Date, project.Description, pq.Array(project.Images), pq.Array(project.Links))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// create user
func CreateUser(c *gin.Context, db *sql.DB) {
	user := models.User{}
	c.BindJSON(&user)

	sqlStatement := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`

	hashedPassword := user.HashPassword(user.Password)

	err := db.QueryRow(sqlStatement, user.Username, user.Email, hashedPassword).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
}

// user login posts
func LoginPost(c *gin.Context, db *sql.DB) {
	user := models.User{}
	session := sessions.Default(c)
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		return
	}

	if helpers.EmptyUserPass(user.Username, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Parameters can't be empty."})
		log.Println("Parameters can't be empty")
		return
	}

	var hashedPassword string

	sqlStatement := `SELECT password FROM users WHERE username = $1;`

	row := db.QueryRow(sqlStatement, user.Username)

	if err := row.Scan(&hashedPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User is not found."})
		log.Println("User is not found.")
		return
	}

	if !helpers.CheckUserPass(user.Password, hashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect username or password."})
		log.Println("Incorrect password.")
		return
	}

	session.Set(globals.Userkey, user.Username)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}

// user login
func Dashboard(c *gin.Context, db *sql.DB) {
	user := models.User{}

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

// Read a user
func ReadUser(c *gin.Context, db *sql.DB) {
	user := models.User{}

	session := sessions.Default(c)
	username := session.Get(globals.Userkey)

	sqlStatement := `SELECT id, username, email FROM users WHERE username = $1;`

	row := db.QueryRow(sqlStatement, username)
	_ = row.Scan(&user.ID, &user.Username, &user.Email)

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}

// update a user
func UpdateUser(c *gin.Context, db *sql.DB) {
	user := models.User{}

	c.BindJSON(&user)
	session := sessions.Default(c)
	username := session.Get(globals.Userkey)

	if helpers.EmptyUserPass(user.Username, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Parameters can't be empty."})
		log.Printf("Parameters can't be empty")
		return
	}

	var hashedPassword string

	sqlStatement := `SELECT password FROM users WHERE username = $1;`

	row := db.QueryRow(sqlStatement, username)

	if err := row.Scan(&hashedPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User is not found."})
		log.Println("User is not found")
		return
	}

	if !helpers.CheckUserPass(user.Password, hashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect username or password."})
		log.Println("Incorrect username or password")
		return
	}

	updateUserSqlStatement := `UPDATE users SET username = $2, email = $3, password = $4 WHERE id = $1;`

	var err error

	if user.NewPassword != "" {
		hashedPassword := user.HashPassword(user.NewPassword)
		_, err = db.Exec(updateUserSqlStatement, user.ID, user.Username, user.Email, hashedPassword)
	} else {
		_, err = db.Exec(updateUserSqlStatement, user.ID, user.Username, user.Email, hashedPassword)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
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
