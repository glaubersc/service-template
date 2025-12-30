package rest

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
)

func OpenAPIMiddleware(doc *openapi3.T) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO:
			// Validate request/response against OpenAPI spec using doc

			next.ServeHTTP(w, r)
		})
	}
}
