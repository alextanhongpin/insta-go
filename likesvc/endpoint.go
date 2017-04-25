package likesvc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alextanhongpin/instago/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/alextanhongpin/instago/helper/httputil"
)

// Endpoint is the struct that holds all the endpoints for the auth service
type Endpoint struct {
	DB *Service
}

func (e Endpoint) Like(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get the userID from the middleware
	userID, ok := r.Context().Value("user_id").(string)
	if !ok || userID == "" {
		httpUtil.Error(w, "You need to be logged in to like the photo", http.StatusForbidden)
		return
	}

	var like Like
	// Decode the body into our like struct
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		httpUtil.Error(w, "Malformed body", http.StatusBadRequest)
		return
	}
	like.UserID = userID

	if err = e.DB.Like(like); err != nil {
		httpUtil.Error(w, "Error liking photo", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, `{"ok":true}`)
}

func (e Endpoint) Unlike(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get the userID from the middleware
	userID, ok := r.Context().Value("user_id").(string)
	if !ok || userID == "" {
		httpUtil.Error(w, "You need to be logged in to like the photo", http.StatusForbidden)
		return
	}
	
	var like Like
	// Decode the body into our like struct
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		httpUtil.Error(w, "Malformed body", http.StatusBadRequest)
		return
	}
	like.UserID = userID

	if err = e.DB.Unlike(like); err != nil {
		httpUtil.Error(w, "Error unliking photo", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, `{"ok":true}`)
}

func (e Endpoint) Count(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	photoID := ps.ByName("photoID")
	count, err := e.DB.Count(Like{PhotoID: photoID})
	if err != nil {
		httpUtil.Error(w, "Error getting photos count", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, `{"like_count": %v, "photo_id": %v }`, count, photoID)
}
