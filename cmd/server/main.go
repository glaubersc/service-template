package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/your-org/service-template/internal/infrastructure/config"
)

func main() {
	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Printf("starting service %s on port %s", cfg.ServiceName, cfg.HTTPPort)

	<-ctx.Done()
	log.Println("shutting down service gracefully")
}
