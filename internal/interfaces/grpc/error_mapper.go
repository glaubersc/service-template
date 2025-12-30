package grpciface

import (
	"github.com/glaubersc/ecosystem/services/service-template/internal/domain/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapError(err error) error {
	if err == nil {
		return nil
	}

	domainErr, ok := err.(service.DomainError)
	if !ok {
		return status.Error(codes.Internal, "internal server error")
	}

	switch domainErr.Code {
	case service.ErrInvalidArgument:
		return status.Error(codes.InvalidArgument, domainErr.Message)
	case service.ErrNotFound:
		return status.Error(codes.NotFound, domainErr.Message)
	case service.ErrFailedPrecondition:
		return status.Error(codes.FailedPrecondition, domainErr.Message)
	case service.ErrConflict:
		return status.Error(codes.Aborted, domainErr.Message)
	case service.ErrDependencyFailure:
		return status.Error(codes.Unavailable, domainErr.Message)
	default:
		return status.Error(codes.Internal, domainErr.Message)
	}
}
