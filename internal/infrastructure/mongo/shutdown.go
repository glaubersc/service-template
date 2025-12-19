package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func Disconnect(ctx context.Context, client *mongo.Client) error {
	return client.Disconnect(ctx)
}
