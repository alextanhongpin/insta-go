package photosvc

import (
	"encoding/json"
	"net/http"

	"github.com/alextanhongpin/instago/helper"
	"github.com/julienschmidt/httprouter"
)

type Endpoint struct{}

func (e Endpoint) All(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := allRequest{
			Query: "",
		}
		v, err := svc.All(req)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}

		res := allResponse{
			Data: v,
		}

		j, err := json.Marshal(res)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}
		helper.ResponseWithJSON(w, j, 200)
	}
}
func (e Endpoint) One(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		req := oneRequest{
			ID: ps.ByName("id"),
		}
		v, err := svc.One(req)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}

		res := oneResponse{
			Data: v,
		}
		j, err := json.Marshal(res)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}
		helper.ResponseWithJSON(w, j, 200)
	}
}

func (e Endpoint) Create(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Insert into the database
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
