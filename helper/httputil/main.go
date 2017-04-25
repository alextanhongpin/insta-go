// httputil.go is a utility that makes it convenient to throw
// http error

package httpUtil

import (
	// "fmt"
	"net/http"
)

// func Error(w http.ResponseWriter, message string, code int) {
// 	w.Header().Set("Content-Type", "application/vnd.api+json; charset=utf-8")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, `{"message": "%v"}`, message)
// }

func Json(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/vnd.api+json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}
