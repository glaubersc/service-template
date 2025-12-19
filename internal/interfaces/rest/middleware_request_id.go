package rest

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	RequestIDHeader     = "X-Request-Id"
	CorrelationIDHeader = "X-Correlation-Id"
)

type ctxKey string

const (
	requestIDKey ctxKey = "requestId"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get(RequestIDHeader)
		if reqID == "" {
			reqID = uuid.NewString()
		}

		corrID := r.Header.Get(CorrelationIDHeader)
		if corrID == "" {
			corrID = reqID
		}

		ctx := context.WithValue(r.Context(), requestIDKey, corrID)

		w.Header().Set(RequestIDHeader, reqID)
		w.Header().Set(CorrelationIDHeader, corrID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TraceIDFromContext(ctx context.Context) string {
	if v, ok := ctx.Value(requestIDKey).(string); ok {
		return v
	}
	return ""
}
