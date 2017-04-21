package authsvc

import (
	"errors"
	"fmt"

	"database/sql"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
)

// Service holds the database context
type Service struct {
	*sql.DB
}

var (
	errorInvalidUser = errors.New("Err not found")
)

// Login checks if the user have access to login
func (s Service) Login(request interface{}) (LoginResponse, error) {
	// Handle the users's login
	var res LoginResponse
	req := request.(LoginRequest)
	if !(req.Email == "john.doe@mail.com" && req.Password == "123456") {
		// Found the correct user
		// Create access token and return it
		return res, errorInvalidUser
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "userprofileid123",
	})
	// Sign and get the complete encoded token as a string
	//  using the secret
	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return res, err
	}
	res.AccessToken = accessToken
	// Check if user's exist in the database
	// Check is the email is unique
	// Hash the password
	// Create a jwt token
	// Returns the user's token
	return res, nil
}

// Register checks if the user can register a new account
func (s Service) Register(request interface{}) (User, error) {
	var user User

	// Expires the token and cookie in 1 hour
	// expireToken := time.Now().Add(time.Hour * 1).Unix()

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
