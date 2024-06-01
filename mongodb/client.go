package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MustClient connect to mongo and return connection or panic on error
func MustClient(ctx context.Context, cfg *Config) *mongo.Client {
	client, err := NewClient(ctx, cfg)
	if err != nil {
		panic(err)
	}
	return client
}

// NewClient from Windmill resource
func NewClient(ctx context.Context, cfg *Config) (*mongo.Client, error) {
	opt := options.Client().ApplyURI(cfg.DSN)
	return mongo.Connect(ctx, opt)
}
