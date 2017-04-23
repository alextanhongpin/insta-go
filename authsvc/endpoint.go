package authsvc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alextanhongpin/instago/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/alextanhongpin/instago/common"
	"github.com/alextanhongpin/instago/helper"
)

// Endpoint is the struct that holds all the endpoints for the auth service
type Endpoint struct{}

const CONTENT_TYPE_VND string = "application/vnd.api+json; charset=utf-8"

var service *Service

func init() {
	// Initialize the service before starting the packages
	service = &Service{common.GetDatabaseContext()}
}

// Login endpoint handles the user login
func (e Endpoint) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()

	//email := r.Form["email"][0]
	//password := r.Form["password"][0]
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := service.GetUserByEmail(GetUserRequest{Email: email})

	if err != nil {
		helper.ErrorWithJSON(w, err.Error(), 400)
		return
	}

	// NO users found
	if (User{}) == user {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{
			"error": "Forbidden request",
			"message": "User with the %v is not found"
			}`, email)
		return
	}

	// Validate if the password is correct
	if passwordMatched := helper.CheckPasswordHash(password, user.Password); !passwordMatched {
		w.Header().Set("Content-Type", CONTENT_TYPE_VND)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{
			"error": "Bad Request",
			"message": "Email or password is incorrect"
		}`)
		return
	}

	expireCookie := time.Now().Add(time.Hour * 1)

	jwtToken, _ := helper.CreateJWTToken(user.ID)

	// Place the token in the client's cookie
	cookie := http.Cookie{
		Name:     "Auth",
		Value:    jwtToken,
		Expires:  expireCookie,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/profile", http.StatusFound)
}

func (e Endpoint) LoginView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token, _ := r.Cookie("Auth")
	// if err != nil {
	// 	http.NotFound(w, r)
	// 	return
	// }
	if token != nil {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}
	helper.RenderTemplate(w, "login", "base", nil)
}

func (e Endpoint) Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	fmt.Println("\nAt /register route")

	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Printf("Registering user with %v and %v\n", password, email)

	// Check if the user exists
	user, err := service.GetUserByEmail(GetUserRequest{Email: email})
	fmt.Printf("User found - %v", user)
	// Assert that the interface is of User type
	// user = user.(GetUserResponse)

	if err != nil {
		w.Header().Set("Content-Type", CONTENT_TYPE_VND)
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{
			"error": "%v"
		}`, err.Error())
		return
	}

	if (User{}) != user {
		if passwordMatched := helper.CheckPasswordHash(password, user.Password); !passwordMatched {
			w.Header().Set("Content-Type", CONTENT_TYPE_VND)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{
			"error": "Bad Request",
			"message": "Email or password is incorrect"
		}`)
			return
		}
		// User found, redirect
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}
	// Create a new user
	// Hash the password first
	hashedPassword, _ := helper.HashPassword(password)
	newuser := User{
		Email:    email,
		Password: hashedPassword,
	}
	userID, err := service.CreateUser(newuser)
	if err != nil {
		w.Header().Set("Content-Type", CONTENT_TYPE_VND)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{
				"error": "Bad Request",
				"message": "%v"
			}`, err.Error())
		return
	}

	expireCookie := time.Now().Add(time.Hour * 1)

	fmt.Println("UserID", userID)
	jwtToken, _ := helper.CreateJWTToken(userID.(string))

	// Place the token in the client's cookie
	cookie := http.Cookie{
		Name:     "Auth",
		Value:    jwtToken,
		Expires:  expireCookie,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	// Create Cookie
	// w.Header().Set("Content-Type", CONTENT_TYPE_VND)
	// fmt.Fprint(w, `{"ok":true}`)
	http.Redirect(w, r, "/profile", http.StatusFound)

}

func (e Endpoint) RegisterView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	helper.RenderTemplate(w, "register", "base", nil)
}

// Profile returns the
func (e Endpoint) Profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		fmt.Println("This user is not valid", userID, ok)
	}

	user, err := service.GetUserByID(GetUserRequest{ID: userID})

	if err != nil {
		helper.ErrorWithJSON(w, err.Error(), 400)
		return
	}
	if (User{}) == user {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{
			"error": "Forbidden request",
			"message": "User with the %v is not found"
			}`, userID)
		return
	}

	// fmt.Println("getting user id", userId)
	helper.RenderTemplate(w, "profile", "base", struct {
		Email, Username, UserID, FirstName, LastName string
	}{user.Email, user.Username, user.ID, user.FirstName, user.LastName})

}

func (e Endpoint) UserView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// id := ps.ByName("id")

}
func (e Endpoint) UsersView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// id := ps.ByName("id")
	users, err := service.GetUsers(nil)
	// fmt.Println("Get users - ", users)
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}
	helper.RenderTemplate(w, "users", "base", users)
}

func (e Endpoint) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deleteCookie := http.Cookie{
		Name:    "Auth",
		Value:   "none",
		Expires: time.Now(),
	}
	http.SetCookie(w, &deleteCookie)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func (e Endpoint) UpdateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", CONTENT_TYPE_VND)

	userID, _ := r.Context().Value("user_id").(string)

	_, err := service.UpdateUser(User{
		ID:        userID,
		Username:  r.FormValue("username"),
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
	})
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "%v"}`, err)
		return
	}

	http.Redirect(w, r, "/profile", http.StatusFound)
}
