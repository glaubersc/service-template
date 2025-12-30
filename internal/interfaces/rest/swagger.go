package rest

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func SwaggerHandler() http.Handler {
	return httpSwagger.Handler(
		httpSwagger.URL("/swagger/openapi.yaml"),
	)
}
