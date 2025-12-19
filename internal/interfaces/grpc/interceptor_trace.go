package grpciface

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	RequestIDHeader     = "x-request-id"
	CorrelationIDHeader = "x-correlation-id"
)

func TraceInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {

	md, _ := metadata.FromIncomingContext(ctx)

	reqID := getOrCreate(md, RequestIDHeader)
	corrID := getOrCreate(md, CorrelationIDHeader)

	newMD := metadata.New(map[string]string{
		RequestIDHeader:     reqID,
		CorrelationIDHeader: corrID,
	})

	ctx = metadata.NewIncomingContext(ctx, newMD)

	return handler(ctx, req)
}

func getOrCreate(md metadata.MD, key string) string {
	if values := md.Get(key); len(values) > 0 {
		return values[0]
	}
	return uuid.NewString()
}
