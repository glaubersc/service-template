package rest

import (
	"net/http"

	"github.com/glaubersc/ecosystem/services/service-template/internal/domain/service"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				writeError(
					w,
					r,
					service.New(
						service.ErrInternal,
						"internal server error",
						nil,
					),
					http.StatusInternalServerError,
				)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
