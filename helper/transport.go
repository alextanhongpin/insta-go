package helper

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// ErrorWithJSON returns the json error object with the appropriate content-type header
func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/vnd.api+json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, `{"message": "%q"}`, message)
	//w.Write([]byte(fmt.Sprintf("{message: %s}", message)))
}

// ResponseWithJSON returns the json resource or collection with the appropriate content-type header
func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/vnd.api+json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

// FetchParams returns the param object from the request context
func FetchParams(r *http.Request) httprouter.Params {
	ctx := r.Context()
	return ctx.Value("params").(httprouter.Params)
}
