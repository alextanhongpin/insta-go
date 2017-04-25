package photosvc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alextanhongpin/instago/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/alextanhongpin/instago/helper/httputil"
)

type Endpoint struct{}

// func (e Endpoint) All(svc *Service) http.HandlerFunc {
func (e Endpoint) All(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		userID, _ := r.Context().Value("user_id").(string)

		req := allRequest{
			UserID: userID,
		}
		v, err := svc.All(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := allResponse{
			Data: v,
		}

		js, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httpUtil.Json(w, js, http.StatusOK)
	}
}
func (e Endpoint) One(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		req := oneRequest{
			ID: ps.ByName("id"),
		}
		v, err := svc.One(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := oneResponse{
			Data: v,
		}
		js, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httpUtil.Json(w, js, http.StatusOK)
	}
}

func (e Endpoint) Create(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		userID, ok := r.Context().Value("user_id").(string)
		if !ok || userID == "" {
			http.Error(w, "No user with the id available", http.StatusForbidden)
			return
		}

		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		caption := r.FormValue("caption")

		img := Image{handler.Filename}
		imgPath, relativeImgPath := img.Path("/static/images/")
		f, err := os.OpenFile(relativeImgPath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer f.Close()
		io.Copy(f, file)

		photoID, err := svc.Create(Photo{
			Src:     imgPath,
			Caption: caption,
			UserID:  userID,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, `{"photo_id": %v}`, photoID)
	}
}

func (e Endpoint) Count(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Update the database
		// Get user ID from context and assert the type
		userID, ok := r.Context().Value("user_id").(string)
		if !ok {
			http.Error(w, "You are not authorized to use this service", http.StatusForbidden)
			return
		}

		count, err := svc.Count(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, `{"data": { "count": %v }}`, count)
	}
}

func (e Endpoint) Update(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Update the database
	}
}

func (e Endpoint) Delete(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Delete an entry
	}
}
