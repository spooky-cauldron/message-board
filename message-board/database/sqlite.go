package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	// "os"
)

const createTableQuery string = `
CREATE TABLE 'messages' (
    'id' UUID,
    'text' VARCHAR(64)
)
`

func InitSqlite() *sql.DB {
	dbSource := ":memory:"
	db, err := sql.Open("sqlite3", dbSource)
	check(err)

	// createTableQuery, err := os.ReadFile("../sql/messages_table.sql")
	// check(err)

	_, err = db.Exec(string(createTableQuery))
	check(err)

	log.Println("Initialized Database.")
	return db
}
