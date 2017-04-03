package photosvc

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	//"github.com/justinas/alice"
	"go.uber.org/zap"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "this is a photo!")
	logger, _ := zap.NewProduction()
	logger.Info("Failed to fetch URL.",
		// Structured context as strongly-typed Field values.
		zap.String("url", "url"),
		zap.Int("attempt", 1),
		zap.Duration("backoff", 1000),
	)
}

// How to carry out chaining with httprouter + alice?
// func MiddlewareOne(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Executing middleware 1")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func MiddlewareTwo(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Executing middleware 2")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func Final(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello world Emd pf middleware")
// }

// func Wrap() httprouter.Handle {
// 	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 		alice.New(MiddlewareOne, MiddlewareTwo).Then(Final)
// 	}
// }

func Something(params string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	}
}

func Init(router *httprouter.Router) {
	//router.GET("/photos", Wrap())
	router.GET("/photos", Index)
}
