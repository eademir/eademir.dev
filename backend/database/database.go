package database

import (
	"database/sql"
	"fmt"

	//"log"
	"os"
	//"github.com/joho/godotenv"
)

var port int
var user, password, dbname, host string

func DatabaseConnection() (db *sql.DB, err error) {
	//if err := godotenv.Load(".env"); err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	port = 5432
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
	host = os.Getenv("DB_HOST")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)

	err = db.Ping()

	return db, err
}
