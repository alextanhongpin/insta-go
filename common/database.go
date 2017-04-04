package common

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

type Photo struct {
	Src     string `db:"src"`
	Caption string `db:"caption"`
}

func InitDatabase() {
	// Disable SSL for development
	db, err := sql.Open("postgres", "user=postgres dbname=instadb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// rows, err := db.Query("SELECT * FROM department WHERE id = $1", 1)
	rows, err := db.Query("SELECT * FROM photo")

	if err != nil {
		fmt.Println(err)
	}

	var photos []Photo
	for rows.Next() {
		// Handle null strings
		var src sql.NullString
		var caption sql.NullString

		err = rows.Scan(&src, &caption)
		if err == nil {
			p := Photo{}
			if src.Valid {
				p.Src = src.String
			}
			if caption.Valid {
				p.Caption = caption.String
			}
			photos = append(photos, p)
		}
	}
	// var photo Photo
	// for i := range rows {
	// 	photo = Json.Unmarshal(&photo)
	// 	photos = append(photos, photo)
	// }
	fmt.Println(photos, len(photos))
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
