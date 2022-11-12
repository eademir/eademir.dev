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
	public.Use(middleware.CORSMiddleware)
	routes.PublicRoutes(public, db)

	private := r.Group("/admin")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private, db)

	r.Run(":8000")
}
