package common

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func init() {
	db = InitDatabase()
}

func GetDatabaseContext() *sql.DB {
	if db != nil {
		return db
	}
	db = InitDatabase()
	return db
}

func InitDatabase() *sql.DB {
	// Disable SSL for development
	// Set parse time to true to enable time.Time  parseTime=True
	db, err = sql.Open("postgres", "user=postgres dbname=instadb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
