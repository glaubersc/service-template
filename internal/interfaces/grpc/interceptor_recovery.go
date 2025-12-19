package grpciface

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {

	defer func() {
		if rec := recover(); rec != nil {
			err = status.Error(codes.Internal, "internal server error")
		}
	}()

	return handler(ctx, req)
}
