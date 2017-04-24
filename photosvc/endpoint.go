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
			httpUtil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := allResponse{
			Data: v,
		}

		js, err := json.Marshal(res)
		if err != nil {
			httpUtil.Error(w, err.Error(), http.StatusBadRequest)
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
			httpUtil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := oneResponse{
			Data: v,
		}
		js, err := json.Marshal(res)
		if err != nil {
			httpUtil.Error(w, err.Error(), http.StatusBadRequest)
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

		caption := r.FormValue("caption")

		fmt.Println("Got caption - ", caption)

		fmt.Fprintf(w, "%v", handler.Header)
		img := Image{handler.Filename}
		imgPath := img.Path("/static/images/")
		fmt.Println("imgPath - ", imgPath)
		f, err := os.OpenFile(imgPath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("error opening file", err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		fmt.Println("The image path is - ", imgPath)
		fmt.Println("The user id is - ", userID)

		photoID, err := svc.Create(Photo{
			Src:     imgPath,
			Caption: caption,
			UserID:  userID,
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
}

func (e Endpoint) Count(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Update the database
		// Get user ID from context and assert the type
		userID, ok := r.Context().Value("user_id").(string)
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, `{"error": "You are not authorized to use this service"}`)
			return
		}
		// Get the number of images that the user has uploaded
		count, err := svc.Count(userID)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"error": "Something happened"}`)
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
