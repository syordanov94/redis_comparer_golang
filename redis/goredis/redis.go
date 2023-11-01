package goredis

import (
	"context"
	"time"

	goRedis "github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client goRedis.Client
}

func NewRedisRepository(address string) RedisRepository {
	return RedisRepository{
		client: *goRedis.NewClient(&goRedis.Options{
			Addr: address,
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

func (repo *RedisRepository) HashGetAll(key string) (map[string]string, error) {
	ctx := context.Background()
	val, err := repo.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (repo *RedisRepository) HashSet(hashKey, key, value string) error {
	ctx := context.Background()
	_, err := repo.client.HSet(ctx, hashKey, key, value).Result()
	if err != nil {
		return err
	}
	return nil
}

func (repo *RedisRepository) ListGet(key string, start, end int) ([]string, error) {
	ctx := context.Background()
	ret, err := repo.client.LRange(ctx, key, int64(start), int64(end)).Result()
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (repo *RedisRepository) ListPush(key string, value string) error {
	ctx := context.Background()
	_, err := repo.client.LPush(ctx, key, value).Result()
	if err != nil {
		return err
	}
	return nil
}
