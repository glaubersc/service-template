package rest

import (
	"encoding/json"
	"net/http"

	"github.com/glaubersc/ecosystem/services/service-template/internal/domain/service"
)

func writeError(w http.ResponseWriter, r *http.Request, err error, status int) {
	traceID := TraceIDFromContext(r.Context())

	domainErr, ok := err.(service.DomainError)
	if !ok {
		domainErr = service.New(
			service.ErrInternal,
			"internal server error",
			nil,
		)
	}

	resp := ErrorResponse{
		Error: ErrorBody{
			Code:    domainErr.Code,
			Message: domainErr.Message,
			Details: domainErr.Details,
			TraceID: traceID,
		},
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
