package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func Logger(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logger, _ := zap.NewProduction()
		logger.Info("Failed to fetch URL.",
			// Structured context as strongly-typed Field values.
			zap.String("url", "url"),
			zap.Int("attempt", 1),
			zap.Duration("backoff", 1000),
		)
		h(w, r, ps)
	}
}
