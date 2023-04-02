package goredis

import (
	"context"
	"rediscomparer/redis"
	"time"

	goRedis "github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client goRedis.Client
}

var _ redis.RedisRepository = (*RedisRepository)(nil)

func NewRedisRepository(address, password string) RedisRepository {
	return RedisRepository{
		client: *goRedis.NewClient(&goRedis.Options{
			Addr:     address,
			Password: password,
		}),
	}
}

func (repo *RedisRepository) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := repo.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (repo *RedisRepository) Set(key, value string) error {
	ctx := context.Background()
	_, err := repo.client.Set(ctx, key, value, 5*time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}
