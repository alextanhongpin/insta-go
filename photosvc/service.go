package photosvc

import (
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"
)

type Service struct {
	*sql.DB
}

func (s Service) All(request interface{}) ([]Photo, error) {
	// req := request.(allRequest)
	rows, err := s.Query("SELECT * FROM photo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		// Handle null strings
		var src sql.NullString
		var caption sql.NullString

		err = rows.Scan(&src, &caption)
		if err != nil {
			return nil, err
		}
		p := Photo{}
		if src.Valid {
			p.Src = src.String
		}
		if caption.Valid {
			p.Caption = caption.String
		}
		photos = append(photos, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return photos, nil
}

func (s Service) One(request interface{}) (Photo, error) {
	req := request.(oneRequest)
	var photo Photo
	var src sql.NullString
	var caption sql.NullString
	err := s.QueryRow("SELECT * FROM photo WHERE src = $1", req.ID).Scan(&src, &caption)

	fmt.Println("error", err)
	if err != nil {
		return photo, err
	}
	if src.Valid {
		photo.Src = src.String
	}
	if caption.Valid {
		photo.Caption = caption.String
	}
	return photo, nil
}
