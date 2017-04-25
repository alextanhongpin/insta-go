package authsvc

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Service holds the database context
type Service struct {
	*sql.DB
}

// All returns the users
func (s Service) All(request interface{}) ([]User, error) {
	var users []User
	rows, err := s.Query(`SELECT username, email, user_id FROM users`)
	defer rows.Close()

	for rows.Next() {
		var username, email, userID sql.NullString

		if err = rows.Scan(&username, &email, &userID); err != nil {
			return nil, err
		}

		var u User
		if username.Valid {
			u.Username = username.String
		}
		if email.Valid {
			u.Email = email.String
		}
		if userID.Valid {
			u.ID = userID.String
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (s Service) GetUserByEmail(req User) (User, error) {
	var res User

	var email sql.NullString
	var password sql.NullString
	var userID sql.NullString

	err := s.QueryRow(`
		SELECT 
			email, 
			salted_password, 
			user_id 
		FROM 
			users 
		WHERE 
			email = $1
	`, req.Email).Scan(&email, &password, &userID)

	switch {
	case err == sql.ErrNoRows:
		return res, nil
	case err != nil:
		return res, err
	}

	if email.Valid {
		res.Email = email.String
	}
	if password.Valid {
		res.Password = password.String
	}
	if userID.Valid {
		res.ID = userID.String
	}
	return res, err
}

func (s Service) GetUserByID(req User) (User, error) {
	var res User

	var email sql.NullString
	var password sql.NullString
	var userID sql.NullString
	var username sql.NullString
	var userphoto sql.NullString
	var firstName sql.NullString
	var lastName sql.NullString

	err := s.QueryRow(`
		SELECT 
			email, 
			salted_password,
			user_id, 
			username,
			userphoto,
			first_name,
			last_name 
		FROM
			users 
		WHERE 
			user_id = $1
	`, req.ID).Scan(
		&email,
		&password,
		&userID,
		&username,
		&userphoto,
		&firstName,
		&lastName)

	switch {
	case err == sql.ErrNoRows:
		return res, nil
	case err != nil:
		return res, err
	}
	if email.Valid {
		res.Email = email.String
	}
	if password.Valid {
		res.Password = password.String
	}
	if userID.Valid {
		res.ID = userID.String
	}
	if firstName.Valid {
		res.FirstName = firstName.String
	}
	if lastName.Valid {
		res.LastName = lastName.String
	}
	if username.Valid {
		res.Username = username.String
	}
	if userphoto.Valid {
		res.Userphoto = userphoto.String
	}
	return res, err
}

func (s Service) Create(req User) (interface{}, error) {
	var userID string

	err := s.QueryRow(`
		INSERT INTO 
			users (username, email, salted_password) 
		VALUES 
			($1, $2, $3) 
		RETURNING 
			user_id
	`, req.Email, req.Email, req.Password).Scan(&userID)

	if err != nil {
		return "", err
	}
	return userID, nil
}

func (s Service) UpdateUser(req User) error {
	_, err := s.Exec(`
		UPDATE 
			users 
		SET 
			username = $1, first_name = $2, last_name = $3 
		WHERE 
			user_id = $4
	`, req.Username, req.FirstName, req.LastName, req.ID)
	return err
}
func (s Service) UploadPhoto(req User) error {
	_, err := s.Exec(`
		UPDATE 
			users 
		SET 
			userphoto =  $1 
		WHERE
			user_id = $2 
		RETURNING
			user_id
	`, req.Userphoto, req.ID)
	return err
}
