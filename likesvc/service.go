package likesvc

import (
	"database/sql"
	// Blank import
	_ "github.com/lib/pq"
)

// Service holds the database context
type Service struct {
	*sql.DB
}

// Register checks if the user can register a new account
func (s Service) Like(req Like) (bool, error) {
	_ = s.QueryRow("INSERT INTO likes (user_id, photo_id) VALUES ($1, $2)", req.UserID, req.PhotoID)
	return true, nil
}

func (s Service) Unlike(req Like) (bool, error) {
	_ = s.QueryRow("DELETE FROM likes WHERE user_id = $1 AND photo_id = $2", req.UserID, req.PhotoID)
	return true, nil
}

func (s Service) Count(req Like) (int64, error) {
	var count sql.NullInt64
	err := s.QueryRow("SELECT count(photo_id) FROM likes WHERE photo_id = $1", req.PhotoID).Scan(&count)

	if err != nil {
		return 0, err
	}

	if count.Valid {
		return count.Int64, nil
	}
	return 0, nil
}

// Get the count of photos for each users group by user id
// select count(user_id), photos.user_id from photos group by user_id;
