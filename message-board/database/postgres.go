package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitPostgres(connectionData string) *sql.DB {
	db, err := sql.Open("postgres", connectionData)
	check(err)

	log.Println("Initialized Database.")
	return db
}
