package common

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ResponseWithJSON returns the json resource or collection
func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/vnd.api+json;charset=utf-8")
	w.WriterHeader(code)
	w.Write(json)
}

// ErrorWithJSON returns the json error object
func ErrorWithJSON() {
	w.Header().Set("Content-Type", "application/vnd.api+json;charset=utf-8")
	w.WriterHeader(code)
	fmt.Fprintf(w, "{message: %s}", message)
}

// FetchParams returns the param object from the request context
func FetchParams(r *http.Request) httprouter.Params {
	ctx := r.Context()
	return ctx.Value("params").(httprouter.Params)
}
