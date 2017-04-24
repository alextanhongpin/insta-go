package authsvc

import (
	"regexp"
	"time"

	"github.com/ventu-io/go-shortid"
)

// User is the model for the user for this app
type User struct {
	ID        string    `json:"id"`         // ID is a unique identifier for the user
	CreatedAt time.Time `json:"created_at"` // CreatedAt is the date when the user join
	UpdatedAt time.Time `json:"updated_at"` // UpdatedAt is the date when the user's profile is last updated
	Username  string    `json:"username"`   // Username is the name displayed by the user
	FirstName string    `json:"first_name"` // FirstName is the user's first name
	LastName  string    `json:"last_name"`  // LastName
	Userphoto string    `json:"userphoto"`  // Userphoto is the photo displayed by the user
	Email     string    `json:"email"`      // Email is created user register
	Password  string    `json:"password"`   // Password is the hashed password that user use to register
}

// Login is the request required when the user log in the application
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest is the request required when the user register a new account
type RegisterRequest struct {
}

type Image struct {
	Name string
}

func (i Image) Path(path string) (string, string) {
	id, _ := shortid.Generate()
	reg := regexp.MustCompile(`\.(gif|jpg|jpeg|tiff|png)$`)

	match := reg.FindStringSubmatch(i.Name)
	ext := match[0]

	res := reg.ReplaceAllString(i.Name, "$1W")
	finalPath := path + res + id + ext

	// Return an absolute path and relative path
	return finalPath, "." + finalPath // unique id
}

// RegisterResponse is the response returned when the user successfully/failed to register
type RegisterResponse struct{}

type GetUserRequest struct {
	Email string
	ID    string
}
