package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const createTableQuery string = `
CREATE TABLE 'messages' (
    'id' UUID,
    'text' VARCHAR(64)
)
`

func InitSqliteMem() *sql.DB {
	dbSource := ":memory:"
	db := InitSqlite(dbSource)
	db.SetMaxOpenConns(1) // prevent race conditions
	return db
}

func InitSqlite(dbSource string) *sql.DB {
	db, err := sql.Open("sqlite3", dbSource)
	check(err)

	_, err = db.Exec(string(createTableQuery))
	check(err)

	log.Println("Initialized Database.")
	return db
}
