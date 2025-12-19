package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/glaubersc/ecosystem/services/service-template/internal/infrastructure/messaging"
	"github.com/glaubersc/ecosystem/services/service-template/internal/infrastructure/mongo"
	"github.com/your-org/service-template/internal/infrastructure/config"
	grpciface "github.com/your-org/service-template/internal/interfaces/grpc"
	"github.com/your-org/service-template/internal/interfaces/rest"
)

func main() {
	cfg := config.Load()

	// NATS
	natsClient, err := messaging.Connect(cfg.NatsURL)
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	defer natsClient.Close()

	// MongoDB
	mongoClient, err := mongo.Connect(ctx, cfg.MongoURI, cfg.MongoDB)
	if err != nil {
		log.Fatalf("mongo connection failed: %v", err)
	}

	// REST
	router := rest.NewRouter()
	httpServer := rest.NewServer(router, cfg.HTTPPort)

	// gRPC
	grpcCore := grpciface.NewGRPCServer()
	grpcServer, err := grpciface.NewServer(grpcCore, cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to start gRPC: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("HTTP running on %s", cfg.HTTPPort)
		httpServer.Start()
	}()

	go func() {
		log.Printf("gRPC running on %s", cfg.GRPCPort)
		grpcServer.Start()
	}()

	<-ctx.Done()
	log.Println("shutdown started")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongo.Disconnect(shutdownCtx, mongoClient.DB.Client())
	httpServer.Shutdown(shutdownCtx)
	grpcServer.Stop()
}
