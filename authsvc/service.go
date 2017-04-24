package authsvc

import (
	"errors"
	"fmt"

	"database/sql"
	"github.com/dgrijalva/jwt-go"
	// Blank import
	_ "github.com/lib/pq"
)

// Service holds the database context
type Service struct {
	*sql.DB
}

var (
	errorInvalidUser = errors.New("Err not found")
)

// Register checks if the user can register a new account
func (s Service) Register(request interface{}) (User, error) {
	var user User

	// NewWithClaims create a new token by specifying signing method
	// and contains claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "USER_ID",
	})

	// Sign and get the complete encoded token as a string
	//  using the secret
	tokenString, err := token.SignedString([]byte("secret"))
	fmt.Println(tokenString, err)

	return user, nil
}

// Profile returns the current user profile
func (s Service) Profile(request interface{}) (User, error) {
	var user User
	// Search for the current user profile
	// If the user does not exists, throw error
	return user, nil
}

func (s Service) GetUsers(request interface{}) ([]User, error) {
	var res []User
	rows, err := s.Query("SELECT username, email, user_id FROM users")
	defer rows.Close()

	for rows.Next() {
		// Handle null strings
		var username sql.NullString
		var email sql.NullString
		var userID sql.NullString

		err = rows.Scan(&username, &email, &userID)
		if err != nil {
			return nil, err
		}
		u := User{}
		if username.Valid {
			u.Username = username.String
		}
		if email.Valid {
			u.Email = email.String
		}
		if userID.Valid {
			u.ID = userID.String
		}
		res = append(res, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func (s Service) GetUserByEmail(request interface{}) (User, error) {
	fmt.Println("At service/GetUser")
	req := request.(GetUserRequest)
	var res User

	var email sql.NullString
	var password sql.NullString
	var userID sql.NullString

	err := s.QueryRow("SELECT email, salted_password, user_id FROM users where email = $1", req.Email).Scan(&email, &password, &userID)

	// switch err {
	// case sql.ErrNoRows:
	// 	return res, nil
	// case nil:
	// 	return res, err
	// }
	if err != nil {
		if err == sql.ErrNoRows {
			// No results found
			return res, nil
		} else {
			return res, err
		}
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

func (s Service) GetUserByID(request interface{}) (User, error) {
	fmt.Println("At service/GetUser")
	req := request.(GetUserRequest)
	var res User

	var email sql.NullString
	var password sql.NullString
	var userID sql.NullString
	var username sql.NullString
	var userphoto sql.NullString
	var firstName sql.NullString
	var lastName sql.NullString

	err := s.QueryRow("SELECT email, salted_password, user_id, username, userphoto, first_name, last_name FROM users where user_id = $1", req.ID).Scan(&email, &password, &userID, &username, &userphoto, &firstName, &lastName)

	if err != nil {
		if err == sql.ErrNoRows {
			// No results found
			return res, nil
		} else {
			return res, err
		}
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

func (s Service) CreateUser(request interface{}) (interface{}, error) {
	req := request.(User)
	var userID string

	err := s.QueryRow("INSERT INTO users (username, email, salted_password) VALUES ($1, $2, $3) RETURNING user_id", req.Email, req.Email, req.Password).Scan(&userID)
	if err != nil {
		return "", errors.New("Error creating user. Please try again in a few minutes")
	}
	return userID, nil
}

func (s Service) UpdateUser(req User) (bool, error) {
	fmt.Printf("updating user - %+v", req)
	var userID sql.NullString
	row := s.QueryRow("UPDATE users SET username = $1, first_name = $2, last_name = $3 WHERE user_id = $4 RETURNING user_id", req.Username, req.FirstName, req.LastName, req.ID).Scan(&userID)

	fmt.Println(row, userID)
	if userID.String == "" {
		return false, errors.New("No user found with the id")
	}

	return true, nil
}
func (s Service) UploadPhoto(req User) (string, error) {
	fmt.Println("at service/upload user photo", req)
	var userID sql.NullString
	err := s.QueryRow("UPDATE users SET userphoto =  $1 WHERE user_id = $2 RETURNING user_id", req.Userphoto, req.ID).Scan(&userID)
	fmt.Println(err)
	if err != nil {
		return "", err
	}
	if userID.Valid {
		fmt.Println("userID is vlaud", userID)
		return userID.String, nil
	}
	return "", nil
}

// switch {
// case err == sql.ErrNoRows:
//         log.Printf("No user with that ID.")
// case err != nil:
//         log.Fatal(err)
// default:
//         fmt.Printf("Username is %s\n", username)
// }
