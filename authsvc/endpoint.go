package authsvc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alextanhongpin/instago/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/alextanhongpin/instago/helper"
)

// Endpoint is the struct that holds all the endpoints for the auth service
type Endpoint struct{}

// Login endpoint handles the user login
func (e Endpoint) Login(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println("at /Login")
		r.ParseForm()
		fmt.Println(r.Form["email"], r.Form["password"])
		req := LoginRequest{
			Email:    r.Form["email"][0],
			Password: r.Form["password"][0],
		}

		// Check if the user exists
		v, err := svc.Login(req)
		fmt.Println(v)
		if err != nil {
			// Redirect to 404 page
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}
		// claims, ok := req.Context().Value(MyKey).(Claims)
		// if !ok {
		//   http.NotFound(w, r)
		// }
		// claims.UserID
		expireCookie := time.Now().Add(time.Hour * 1)
		// Place the token in the client's cookie
		fmt.Println("AccessToken", v.AccessToken)
		cookie := http.Cookie{
			Name:     "Auth",
			Value:    v.AccessToken,
			Expires:  expireCookie,
			HttpOnly: true,
		}

		fmt.Println("Auth:cookie", v)
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/me", http.StatusFound)
	}
}

func (e Endpoint) Register(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {}
}

// Profile returns the
func (e Endpoint) Profile(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Calling the next value here
		fmt.Println("Calling the next line")
		// ctx := context.WithValue(r.Context(), "query", queryValues)
		// r = r.WithContext(ctx)

		// v := r.Context().Value("params").(url.Values)
		userID, ok := r.Context().Value("user_id").(string)
		if !ok {
			fmt.Println("This user is not valid", userID, ok)
		}
		// fmt.Println("getting user id", userId)
		helper.RenderTemplate(w, "profile", "base", struct {
			Name, UserID string
		}{"john.doe", userID})
	}
}

func (e Endpoint) LoginView(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		helper.RenderTemplate(w, "login", "base", nil)
	}
}

func (e Endpoint) RegisterView(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		helper.RenderTemplate(w, "register", "base", nil)
	}
}

func (e Endpoint) UserView(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {}
}

func (e Endpoint) Logout(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		deleteCookie := http.Cookie{
			Name:    "Auth",
			Value:   "none",
			Expires: time.Now(),
		}
		http.SetCookie(w, &deleteCookie)
		return
	}
}
