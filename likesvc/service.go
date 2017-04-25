// likesvc/service.go contains the methods to carry out CRUD operations for the `likes` resources 
package likesvc

import (
	"database/sql"
	// Blank import
	_ "github.com/lib/pq"
)

// Service holds the db context
type Service struct {
	*sql.DB
}

// Like creates a new entry
func (s Service) Like(req Like) error {
	_, err := s.Exec("INSERT INTO likes (user_id, photo_id) VALUES ($1, $2)", req.UserID, req.PhotoID)
	return err
}

// Unlike remove the entry
func (s Service) Unlike(req Like) error {
	_, err := s.Exec("DELETE FROM likes WHERE user_id = $1 AND photo_id = $2", req.UserID, req.PhotoID)
	return err
}

// Count returns the like count for a photo
func (s Service) Count(req Like) (int64, error) {
	var count sql.NullInt64
	if err := s.QueryRow("SELECT count(photo_id) FROM likes WHERE photo_id = $1", req.PhotoID).Scan(&count); err != nil {
		return 0, err
	}

	if count.Valid {
		return count.Int64, nil
	}
	return 0, nil
}
