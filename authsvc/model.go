package authsvc

import (
	"time"
)

// User is the model for the user for this app
type User struct {
	ID        string    `json:"id"`         // ID is a unique identifier for the user
	CreatedAt time.Time `json:"created_at"` // CreatedAt is the date when the user join
	UpdatedAt time.Time `json:"updated_at"` // UpdatedAt is the date when the user's profile is last updated
	Username  string    `json:"username"`   // Username is the name displayed by the user
	Userphoto string    `json:"userphoto"`  // Userphoto is the photo displayed by the user
	Email     string    `json:"email"`      // Email is created user register
	Password  string    `json:"password"`   // Password is the hashed password that user use to register
}

// Login is the request required when the user log in the application
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

// RegisterRequest is the request required when the user register a new account
type RegisterRequest struct {
}

// RegisterResponse is the response returned when the user successfully/failed to register
type RegisterResponse struct{}
