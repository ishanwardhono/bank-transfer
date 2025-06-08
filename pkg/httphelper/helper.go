package httphelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	apperrors "github.com/ishanwardhono/transfer-system/pkg/errors"
	"github.com/ishanwardhono/transfer-system/pkg/logger"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func HandleResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

func HandleCreatedResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusCreated)
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

func HandleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	status := http.StatusInternalServerError

	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		status = appErr.Code
	}

	w.WriteHeader(status)
	resp := ErrorResponse{
		Error:   http.StatusText(status),
		Message: err.Error(),
	}
	json.NewEncoder(w).Encode(resp)
}

func LoggerMiddleware(next http.Handler) http.Handler {
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
