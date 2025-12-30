package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(RequestIDMiddleware)
	r.Use(RecoveryMiddleware)
	r.Use(ContentTypeMiddleware)

	r.Get("/health", healthHandler)
	r.Get("/ready", readinessHandler)

	r.Get("/swagger/*", func(w http.ResponseWriter, r *http.Request) {
		SwaggerHandler().ServeHTTP(w, r)
	})

	r.Get("/swagger/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/openapi/openapi.yaml")
	})

	return r
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func readinessHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ready"))
}
