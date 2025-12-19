package repository

import "context"

type ExampleEntity struct {
	ID   string
	Name string
}

type ExampleRepository interface {
	Save(ctx context.Context, e ExampleEntity) error
	FindByID(ctx context.Context, id string) (*ExampleEntity, error)
}
