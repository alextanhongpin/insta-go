package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alextanhongpin/instago/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/alextanhongpin/instago/common"
	"github.com/dgrijalva/jwt-go"
)

// Protect is a middleware that checks if the user is
// authorized to access an endpoint
func Protect(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// If no Auth cookie is set then return a 404 not found
		cookie, err := r.Cookie("Auth")
		if err != nil {
			http.NotFound(w, r)
			return
		}

		// Returns a Token using the cookie
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret,
			// e.g. []byte("my secret key")
			return []byte(common.Config.JWTSecret), nil
		})

		if err != nil {
			http.NotFound(w, r)
			return
		}

		// type claimsContextKey string

		// Grab the token claims and pass it into the original request
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
			r := r.WithContext(ctx)
			h(w, r, ps)
		} else {
			http.NotFound(w, r)
			return
		}
	}
}
