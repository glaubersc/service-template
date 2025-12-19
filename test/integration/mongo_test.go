package integration

import (
	"context"
	"testing"
	"time"

	"github.com/your-org/service-template/internal/infrastructure/mongo"
)

func TestMongoConnection(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, "mongodb://localhost:27017", "test_db")
	if err != nil {
		t.Fatalf("failed to connect to mongo: %v", err)
	}

	err = mongo.Disconnect(ctx, client.DB.Client())
	if err != nil {
		t.Fatalf("failed to disconnect mongo: %v", err)
	}
}
