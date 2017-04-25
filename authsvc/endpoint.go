package authsvc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/alextanhongpin/instago/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/alextanhongpin/instago/helper"
)

// Endpoint struct hold the db context and route method
type Endpoint struct {
	DB *Service
}

// Login is a service that authenticates the user
func (e Endpoint) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := e.DB.GetUserByEmail(User{Email: email})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user is registered in the db
	noUserFound := (User{}) == user
	if noUserFound {
		errorMessage := fmt.Sprintf("The email %v does not exist", email)
		http.Error(w, errorMessage, http.StatusForbidden)
		return
	}

	// Found user, check if the password matches
	if isMatchingPassword := helper.CheckPasswordHash(password, user.Password); !isMatchingPassword {
		http.Error(w, "Email or password is incorrect", http.StatusBadRequest)
		return
	}

	// Email and password matches, create an access token
	jwtToken, err := helper.CreateJWTToken(user.ID)
	if err != nil {
		http.Error(w, "Error generating access token", http.StatusInternalServerError)
		return
	}

	// Create a cookie that stores the access token
	cookie := http.Cookie{
		Name:     "Auth",
		Value:    jwtToken,
		Expires:  time.Now().Add(time.Hour * 1),
		HttpOnly: true,
	}

	// Set the cookie for client
	http.SetCookie(w, &cookie)

	// Redirect the user to profile page after successfully logging in
	http.Redirect(w, r, "/profile", http.StatusFound)
}

// LoginView is the page for the user to log in
func (e Endpoint) LoginView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check the cookie to check if the user is already authenticated
	token, _ := r.Cookie("Auth")

	// Token exists, block them from accessing the login page
	if token != nil {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	// User is not logged in, display the login page
	helper.RenderTemplate(w, "login", "base", nil)
}

// Register is the service that allow user creation
func (e Endpoint) Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Check if the user exists
	user, err := e.DB.GetUserByEmail(User{Email: email})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// Check if the user exist (interface should not be empty struct)
	hasUser := (User{}) != user
	if hasUser {
		// Email exists, check if the password is correct
		if isMatchingPassword := helper.CheckPasswordHash(password, user.Password); !isMatchingPassword {
			// Password is incorrect
			http.Error(w, "Email or password is incorrect", http.StatusForbidden)
			return
		}

		// Email and password is correct, redirect to profile page
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	newUser := User{
		Email:    email,
		Password: hashedPassword,
	}

	userID, err := e.DB.Create(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jwtToken, err := helper.CreateJWTToken(userID.(string))
	if err != nil {
		http.Error(w, "Error generating access token", http.StatusInternalServerError)
		return
	}

	// Place the token in the client's cookie
	cookie := http.Cookie{
		Name:     "Auth",
		Value:    jwtToken,
		Expires:  time.Now().Add(time.Hour * 1),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/profile", http.StatusFound)
}

// RegisterView
func (e Endpoint) RegisterView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	helper.RenderTemplate(w, "register", "base", nil)
}

// Profile
func (e Endpoint) Profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("/profile")
	userID, ok := r.Context().Value("user_id").(string)
	fmt.Println("/profile", userID)
	if !ok {
		http.Error(w, "You are not authorized", http.StatusForbidden)
		return
	}

	user, err := e.DB.GetUserByID(User{ID: userID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	noUserFound := (User{}) == user
	if noUserFound {
		http.Error(w, "User not found", http.StatusForbidden)
		return
	}

	helper.RenderTemplate(w, "profile", "base", struct {
		Email, Username, Userphoto, UserID, FirstName, LastName string
	}{user.Email, user.Username, user.Userphoto, user.ID, user.FirstName, user.LastName})

}

func (e Endpoint) UserView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	user, err := e.DB.GetUserByID(User{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	noUserFound := (User{}) == user
	if noUserFound {
		http.Redirect(w, r, "/users", http.StatusFound)
		return
	}
	helper.RenderTemplate(w, "user", "base", user)
}

func (e Endpoint) UsersView(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := e.DB.All(nil)
	if err != nil {
		panic(err)
	}
	helper.RenderTemplate(w, "users", "base", users)
}

func (e Endpoint) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	clearCookie := http.Cookie{
		Name:    "Auth",
		Value:   "none",
		Expires: time.Now(),
	}
	http.SetCookie(w, &clearCookie)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func (e Endpoint) UpdateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "One or more field provided is invalid", http.StatusBadRequest)
		return
	}

	user.ID, _ = r.Context().Value("user_id").(string)

	if err = e.DB.UpdateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/profile", http.StatusFound)
}

func (e Endpoint) UploadPhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	userID, ok := r.Context().Value("user_id").(string)
	if !ok || userID == "" {
		http.Error(w, "Forbidden access", http.StatusForbidden)
		return
	}

	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	// Insert into the database
	// Upload photo
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	img := Image{handler.Filename}
	imgPath, imgRelativePath := img.Path("/static/resources/users/")

	f, err := os.OpenFile(imgRelativePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	if err := e.DB.UploadPhoto(User{
		Userphoto: imgPath,
		ID:        userID,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, `{"ok": true}`)
}
