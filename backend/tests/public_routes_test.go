package test

import (
	"fmt"
	"testing"

	"os"

	"eademir.dev/controllers"
	"eademir.dev/database"
	"eademir.dev/middleware"
	"eademir.dev/routes"
)

func TestMain(m *testing.M) {
	code, err := run(m)
    if err != nil {
        fmt.Println(err)
    }
    os.Exit(code)
}

func run (m *testing.M) (code int, err error) {
	db, err := database.DatabaseConnection()
	r := controllers.SetupRouter()
	defer db.Close()

	public := r.Group("/")
	routes.PublicRoutes(public, db)

	private := r.Group("/admin")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private, db)

	return m.Run(), err
}
