package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitPostgresConn(connectionData string) *sql.DB {
	db, err := sql.Open("postgres", connectionData)
	check(err)

	log.Println("Initialized Database.")
	return db
}

func InitPostgres() *sql.DB {
	db_host := getEnv("DB_HOST", "localhost")
	db_port := getEnv("DB_PORT", "5432")
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_ssl := getEnv("DB_SSL", "disable")
	connectionData := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		db_host,
		db_port,
		db_name,
		db_user,
		db_pass,
		db_ssl,
	)

	return InitPostgresConn(connectionData)
}
