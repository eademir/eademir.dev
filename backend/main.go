package main

import (
	controllers "eademir.dev/controllers"
	database "eademir.dev/database"
	"eademir.dev/middleware"
	routes "eademir.dev/routes"
	_ "github.com/lib/pq"
)

func main() {
	r := controllers.SetupRouter()

	db, err := database.DatabaseConnection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	public := r.Group("/")
	routes.PublicRoutes(public, db)

	private := r.Group("/api/v1")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private, db)

	r.Run()
}
