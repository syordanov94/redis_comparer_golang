package main

import (
	"fmt"
	"rediscomparer/redis"
	"rediscomparer/redis/redisconfig"
	"rediscomparer/redis/redisgo"
)

func main() {
	// --- (1) ----
	// Get the redis config and init the repository
	config, err := redisconfig.ConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}
	redisRepo := redisgo.NewRedisRepository(fmt.Sprintf("%s:%d", config.Host, config.Port), config.Pass)
	//redisRepo := goredis.NewRepository(fmt.Sprintf("%s:%d", config.Host, config.Port), config.Pass)

	// --- (2) ----
	// Perform the redis operations
	result, err := SetAndGet(&redisRepo, "testKey", "testValue")
	if err != nil {
		panic(err)
	}
	print(result)
}

func SetAndGet(redisRepository redis.RedisRepository, key, value string) (string, error) {

	// --- (1) ---
	// Set the new value into the redis
	err := redisRepository.Set(key, value)
	if err != nil {
		return "", err
	}

	// --- (2) ---
	// Get the new value from redis
	ret, err := redisRepository.Get(key)
	if err != nil {
		return "", err
	}

	return ret, nil
}
