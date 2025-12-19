package mongo

import (
	"context"

	"github.com/your-org/service-template/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type ExampleRepository struct {
	collection string
	client     *Client
}

func NewExampleRepository(client *Client) *ExampleRepository {
	return &ExampleRepository{
		client:     client,
		collection: "examples",
	}
}

func (r *ExampleRepository) Save(ctx context.Context, e repository.ExampleEntity) error {
	_, err := r.client.DB.Collection(r.collection).InsertOne(ctx, e)
	return err
}

func (r *ExampleRepository) FindByID(ctx context.Context, id string) (*repository.ExampleEntity, error) {
	var result repository.ExampleEntity

	err := r.client.DB.
		Collection(r.collection).
		FindOne(ctx, bson.M{"id": id}).
		Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
