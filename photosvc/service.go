package photosvc

import (
	"time"

	"database/sql"
	"github.com/alextanhongpin/instago/helper/pgutil"
	_ "github.com/lib/pq"
)

type Service struct {
	*sql.DB
}

func (s Service) All(req allRequest) ([]Photo, error) {
	rows, err := s.Query(`
		SELECT p.photo_id, p.src, p.caption, p.date_created, count(l.user_id) AS like_count, p.user_id, unnest(array_agg(distinct u.username)) AS username, array_agg(u2.username) AS users
		FROM photos p
			LEFT JOIN users u using(user_id)
				LEFT JOIN likes l ON l.photo_id = p.photo_id
				LEFT JOIN users u2 ON u2.user_id = l.user_id
		WHERE p.user_id = $1
		GROUP BY p.photo_id
	`, req.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		// Handle null strings
		var photoID sql.NullString
		var src sql.NullString
		var caption sql.NullString
		var dateCreated time.Time
		var likeCount sql.NullInt64
		var userID sql.NullString
		var username sql.NullString
		var usersTemp string
		// var users []string

		err = rows.Scan(&photoID, &src, &caption, &dateCreated, &likeCount, &userID, &username, &usersTemp)
		if err != nil {
			return nil, err
		}
		p := Photo{}
		if photoID.Valid {
			p.ID = photoID.String
		}
		if src.Valid {
			p.Src = src.String
		}
		if caption.Valid {
			p.Caption = caption.String
		}
		if likeCount.Valid {
			p.LikeCount = likeCount.Int64
		}
		if userID.Valid {
			p.User.UserID = userID.String
		}
		if username.Valid {
			p.User.Username = username.String
		}
		// Tricky part of converting the postgres array to golang's slice
		p.UserLikes = pgUtil.ArrayToSlice(usersTemp)
		p.CreatedAt = dateCreated
		photos = append(photos, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return photos, nil
}

func B2S(bs []uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}
func (s Service) One(request interface{}) (Photo, error) {
	req := request.(oneRequest)
	var photo Photo
	var src sql.NullString
	var caption sql.NullString

	err := s.QueryRow("SELECT * FROM photo WHERE src = $1", req.ID).Scan(&src, &caption)

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

func (s Service) Count(userID string) (int64, error) {
	var count sql.NullInt64
	err := s.QueryRow("SELECT count(*) FROM photos WHERE user_id = $1", userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	if count.Valid {
		return count.Int64, nil
	}
	return 0, nil
}

func (s Service) Create(req Photo) (string, error) {
	var photoID sql.NullString
	err := s.QueryRow("INSERT INTO photos (src, caption, user_id) VALUES ($1, $2, $3) RETURNING photo_id", req.Src, req.Caption, req.UserID).Scan(&photoID)
	if err != nil {
		return "", err
	}
	if photoID.Valid {
		return photoID.String, nil
	}
	return "", nil
}

// SELECT p.photo_id, p.src, p.caption, p.date_created, count(l.user_id) AS like_count, p.user_id, unnest(array_agg(distinct u.username)) AS username, array_to_json(u2.username) AS users
// FROM photos p
// 	LEFT JOIN users u using(user_id)
// 		LEFT JOIN likes l ON l.photo_id = p.photo_id
// 		LEFT JOIN users u2 ON u2.user_id = l.user_id
// WHERE p.user_id = '2'
// GROUP BY p.photo_id
