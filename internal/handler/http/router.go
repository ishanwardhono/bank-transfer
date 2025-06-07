package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ishanwardhono/transfer-system/pkg/logger"
)

func SetupRouter(router *chi.Mux) {
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(loggerMiddleware)

	router.Get("/", testFunc)
}

// ginLoggerMiddleware creates a custom logger middleware
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL.Path

		// Create a wrapper for the response writer to capture the status code
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		// Process request
		next.ServeHTTP(ww, r)

		// build log format
		msg := fmt.Sprintf("Request: %s %s - %d | %s | %s",
			r.Method,
			path,
			ww.Status(),
			time.Since(start),
			r.RemoteAddr,
		)

		ctx := r.Context()

		if ww.Status() >= 500 {
			logger.Error(ctx, msg)
		} else if ww.Status() >= 400 {
			logger.Warn(ctx, msg)
		} else {
			logger.Info(ctx, msg)
		}
	})
}
