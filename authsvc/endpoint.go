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
	"github.com/alextanhongpin/instago/helper/httputil"
)

// Endpoint is the struct that holds all the endpoints for the auth service
type Endpoint struct {
	DB *Service
}

func (e Endpoint) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := e.DB.GetUserByEmail(GetUserRequest{Email: email})
	if err != nil {
		httpUtil.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If the user struct is empty, it means no user is found
	if (User{}) == user {
		errorMessage := fmt.Sprintf("The email %v does not exist", email)
		httpUtil.Error(w, errorMessage, http.StatusForbidden)
		return
	}

	// Validate if the password is correct
	if passwordMatched := helper.CheckPasswordHash(password, user.Password); !passwordMatched {
		httpUtil.Error(w, "Email or password is incorrect", http.StatusBadRequest)
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
	if token != nil {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}
	helper.RenderTemplate(w, "login", "base", nil)
}

func (e Endpoint) Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Check if the user exists
	user, err := e.DB.GetUserByEmail(GetUserRequest{Email: email})
	if err != nil {
		httpUtil.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	if (User{}) != user {
		if passwordMatched := helper.CheckPasswordHash(password, user.Password); !passwordMatched {
			httpUtil.Error(w, "Email or password is incorrect", http.StatusForbidden)
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
	userID, err := e.DB.CreateUser(newuser)
	if err != nil {
		httpUtil.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expireCookie := time.Now().Add(time.Hour * 1)
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

	user, err := e.DB.GetUserByID(GetUserRequest{ID: userID})

	if err != nil {
		helper.ErrorWithJSON(w, err.Error(), 400)
		return
	}
	if (User{}) == user {
		httpUtil.Error(w, "User not found", http.StatusForbidden)
		return
	}

	// fmt.Println("getting user id", userId)
	helper.RenderTemplate(w, "profile", "base", struct {
		Email, Username, Userphoto, UserID, FirstName, LastName string
	}{user.Email, user.Username, user.Userphoto, user.ID, user.FirstName, user.LastName})

}

func (e Endpoint) UserView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("id")
	fmt.Println("the id is - ", id)

	user, err := e.DB.GetUserByID(GetUserRequest{ID: id})

	// No user found
	if (User{}) == user {
		http.Redirect(w, r, "/users", http.StatusFound)
		return
	}
	if err != nil {
		// httpUtil.Error()
		panic(err)
	}
	helper.RenderTemplate(w, "user", "base", user)
}
func (e Endpoint) UsersView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// id := ps.ByName("id")
	users, err := e.DB.GetUsers(nil)
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
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		httpUtil.Error(w, "One or more field provided is invalid", http.StatusBadRequest)
		return
	}

	user.ID, _ = r.Context().Value("user_id").(string)

	_, err = e.DB.UpdateUser(user)
	if err != nil {
		httpUtil.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/profile", http.StatusFound)
}

func (e Endpoint) UploadPhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	fmt.Println("At upload user photos")
	userID, ok := r.Context().Value("user_id").(string)
	if !ok || userID == "" {
		fmt.Println("Error getting context - ")
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"error": "forbidden request", "message": "%v"}`, "No user id available")
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
		fmt.Println("formFile", err, file, handler)
		return
	}
	defer file.Close()
	// Check if the directory exists, if not create it
	// path := "./static/resources/user/"
	// if _, err := os.Stat(path); os.IsNotExist(err) {

	// 	fmt.Println("Directory does not exists, creating one now")
	// 	os.MkdirAll(path, os.ModePerm)
	// }
	// os.MkdirAll(path, os.ModePerm)

	img := Image{handler.Filename}
	imgPath, imgRelativePath := img.Path("/static/resources/users/")
	fmt.Println("imgPath - ", imgPath)
	f, err := os.OpenFile(imgRelativePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	fmt.Println("The image path is - ", imgPath)
	fmt.Println("The user id is - ", userID)

	photoID, err := e.DB.UploadPhoto(User{
		Userphoto: imgPath,
		ID:        userID,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "bad request", "message": "%v"}`, err.Error())
		return
	}
	fmt.Println("successfully create photo")
	// w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"photo_id": %v}`, photoID)
}
