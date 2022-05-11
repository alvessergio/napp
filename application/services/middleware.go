package services

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func traceIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		traceID := uuid.NewV4().String()

		rw.Header().Set("X-Trace-ID", traceID)
		next.ServeHTTP(rw, r.WithContext(context.WithValue(r.Context(), "traceID", traceID)))
	})
}
