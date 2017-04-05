package common

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() *sql.DB {
	var err error
	fmt.Println(db)
	if db != nil {
		return db
	}
	// Disable SSL for development
	db, err = sql.Open("postgres", "user=postgres dbname=instadb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// func getPhoto() {

// }

// func createPhoto() {
// 	var photo Photo
// 	err := db.QueryRow("INSERT INTO photo(src, caption) VALUES ($1, $2) RETURNING ID", "src", "caption")
// 	if err != nil {
// 		return 0, err
// 	}

// }
