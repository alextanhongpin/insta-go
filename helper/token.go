// token.go is a program that creates a new jwt token from a given
// userID string

package helper

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/alextanhongpin/instago/common"
)

func CreateJWTToken(userID string) (string, error) {
	var jwtToken string
	var err error
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})
	// Sign and get the complete encoded token as a string
	//  using the secret
	jwtToken, err = token.SignedString([]byte(common.Config.JWTSecret))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
