package photosvc

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"go.uber.org/zap"

	"github.com/alextanhongpin/instago/common"
	"github.com/alextanhongpin/instago/middleware"
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
func MiddlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware 1")
		next.ServeHTTP(w, r)
	})
}

func MiddlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware 2")
		next.ServeHTTP(w, r)
	})
}

func Final(w http.ResponseWriter, r *http.Request) {

	queryValues := r.URL.Query()
	ctx := context.WithValue(r.Context(), "query", queryValues)
	r = r.WithContext(ctx)

	v := r.Context().Value("params").(url.Values)
	fmt.Println(v.Get("name"))
	fmt.Fprintf(w, "Hello world Emd pf middleware")
}

func Init(router *httprouter.Router) *httprouter.Router {
	endpoint := Endpoint{}
	service := &Service{common.InitDatabase()}

	// Example chaining middlewares with Alice
	router.Handler("GET", "/photos", alice.New(MiddlewareOne, MiddlewareTwo).ThenFunc(endpoint.All(service)))

	// Example middleware with httprouter
	router.GET("/photos/:id", middleware.Logger(endpoint.One(service)))
	return router
}
