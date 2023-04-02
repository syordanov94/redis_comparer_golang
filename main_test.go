package main

import (
	"fmt"
	"rediscomparer/redis/goredis"
	"rediscomparer/redis/redisconfig"
	"rediscomparer/redis/redisgo"
	"testing"
)

// -------- RedisGo Benchmarks -------------

func BenchmarkRedisGoGetAndSet(b *testing.B) {
	// --- (1) ----
	// Get the redis config and init the repository
	config, err := redisconfig.ConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}
	redisRepo := redisgo.NewRedisRepository(fmt.Sprintf("%s:%d", config.Host, config.Port), config.Pass)

	for i := 0; i < b.N; i++ {
		SetAndGet(&redisRepo, "testKey", "testValue")
	}
}

func BenchmarkRedisGoGet(b *testing.B) {
	// --- (1) ----
	// Get the redis config and init the repository
	config, err := redisconfig.ConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}
	redisRepo := redisgo.NewRedisRepository(fmt.Sprintf("%s:%d", config.Host, config.Port), config.Pass)

	for i := 0; i < b.N; i++ {
		redisRepo.Get("testKey")
	}
}

func BenchmarkRedisGoSet(b *testing.B) {
	// --- (1) ----
	// Get the redis config and init the repository
	config, err := redisconfig.ConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}
	redisRepo := redisgo.NewRedisRepository(fmt.Sprintf("%s:%d", config.Host, config.Port), config.Pass)

	for i := 0; i < b.N; i++ {
		redisRepo.Set("testKey", "testValue")
	}
}

// -----------------------------------------------

// -------- GoRedis Benchmarks -------------------

func BenchmarkGoRedisGetAndSet(b *testing.B) {
	// --- (1) ----
	// Get the redis config and init the repository
	config, err := redisconfig.ConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}
	redisRepo := goredis.NewRedisRepository(fmt.Sprintf("%s:%d", config.Host, config.Port), config.Pass)

	for i := 0; i < b.N; i++ {
		SetAndGet(&redisRepo, "testKey", "testValue")
	}
}

func BenchmarkGoRedisGet(b *testing.B) {
	// --- (1) ----
	// Get the redis config and init the repository
	config, err := redisconfig.ConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}
	redisRepo := goredis.NewRedisRepository(fmt.Sprintf("%s:%d", config.Host, config.Port), config.Pass)

	for i := 0; i < b.N; i++ {
		redisRepo.Get("testKey")
	}
}

func BenchmarkGoRedisSet(b *testing.B) {
	// --- (1) ----
	// Get the redis config and init the repository
	config, err := redisconfig.ConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}
	redisRepo := goredis.NewRedisRepository(fmt.Sprintf("%s:%d", config.Host, config.Port), config.Pass)

	for i := 0; i < b.N; i++ {
		redisRepo.Set("testKey", "testValue")
	}
}

// -----------------------------------------------
