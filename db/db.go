package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	connection := os.Getenv("CONNECTION_DATABASE")
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
