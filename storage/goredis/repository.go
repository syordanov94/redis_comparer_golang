package goredis

import (
	"context"
	"time"

	"rediscomparer/storage"

	goredis "github.com/redis/go-redis/v9"
)

type Repository struct {
	client *goredis.Client
}

var _ storage.Repository = (*Repository)(nil)

func NewRepository(address string) *Repository {
	return &Repository{
		client: goredis.NewClient(&goredis.Options{
			Addr: address,
		}),
	}
}

func (r *Repository) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *Repository) Set(ctx context.Context, key, value string) error {
	return r.client.Set(ctx, key, value, 5*time.Minute).Err()
}
