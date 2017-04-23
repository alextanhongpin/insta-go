// password.go is a helper that hash a password
// and compare if the hash password matches

package helper

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes in a password and converts it to hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash takes in a password and hash string
// and returns a bool to indicate if they match
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
