package grpciface

import "google.golang.org/grpc"

func NewGRPCServer() *grpc.Server {
	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			TraceInterceptor,
			RecoveryInterceptor,
		),
	)
}
