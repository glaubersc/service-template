package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/glaubersc/ecosystem/services/service-template/internal/infrastructure/config"
	"github.com/glaubersc/ecosystem/services/service-template/internal/infrastructure/messaging"
	"github.com/glaubersc/ecosystem/services/service-template/internal/infrastructure/mongo"
	grpciface "github.com/glaubersc/ecosystem/services/service-template/internal/interfaces/grpc"
	"github.com/glaubersc/ecosystem/services/service-template/internal/interfaces/rest"
)

func main() {
	cfg := config.Load()

	// Base context (used for startup only)
	baseCtx := context.Background()

	// --------------------
	// Infrastructure
	// --------------------

	// NATS
	natsClient, err := messaging.Connect(cfg.NatsURL)
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	defer natsClient.Close()

	// MongoDB
	mongoClient, err := mongo.Connect(baseCtx, cfg.MongoURI, cfg.MongoDB)
	if err != nil {
		log.Fatalf("mongo connection failed: %v", err)
	}

	// --------------------
	// Interfaces
	// --------------------

	// REST
	router := rest.NewRouter()
	httpServer := rest.NewServer(router, cfg.HTTPPort)

	// gRPC
	grpcCore := grpciface.NewGRPCServer()
	grpcServer, err := grpciface.NewServer(grpcCore, cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to start gRPC: %v", err)
	}

	// --------------------
	// Signal handling
	// --------------------

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	// --------------------
	// Start servers
	// --------------------

	go func() {
		log.Printf("HTTP running on %s", cfg.HTTPPort)
		if err := httpServer.Start(); err != nil {
			log.Printf("HTTP server stopped: %v", err)
		}
	}()

	go func() {
		log.Printf("gRPC running on %s", cfg.GRPCPort)
		if err := grpcServer.Start(); err != nil {
			log.Printf("gRPC server stopped: %v", err)
		}
	}()

	// --------------------
	// Graceful shutdown
	// --------------------

	<-ctx.Done()
	log.Println("shutdown started")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Order matters
	httpServer.Shutdown(shutdownCtx)
	grpcServer.Stop()
	mongo.Disconnect(shutdownCtx, mongoClient.DB.Client())

	log.Println("shutdown completed")
}
