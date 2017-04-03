package photosvc

import (
	"encoding/json"
	"net/http"
)

type Endpoint struct{}

func (e Endpoint) All(svc service) http.HandlerFunc {
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
