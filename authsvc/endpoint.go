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
		req := LoginRequest{
			Email:    "",
			Password: "",
		}
		v, err := svc.Login(req)
		fmt.Println(v)
		if err != nil {
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
		cookie := http.Cookie{
			Name:     "Auth",
			Value:    "accessToken",
			Expires:  expireCookie,
			HttpOnly: true,
		}

		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/profile", 307)
		// j, err := json.Marshal(res)
		// if err != nil {
		//  helper.ErrorWithJSON(w, err.Error(), 400)
		//  return
		// }
		// helper.ResponseWithJSON(w, j, 200)
	}
}

func (e Endpoint) Register(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	}
}

func (e Endpoint) Profile(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		helper.RenderTemplate(w, "profile", "base", struct{ Name string }{"john.doe"})
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
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	}
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
