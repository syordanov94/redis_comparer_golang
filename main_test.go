package main

import (
	"context"
	"testing"

	"rediscomparer/storage"
	"rediscomparer/storage/goredis"
	"rediscomparer/storage/redigo"

	"github.com/stretchr/testify/require"
)

// -------- Redigo Benchmarks -------------

func TestRedigo(t *testing.T) {
	repository, err := redigo.NewRepository("redis1:6379")
	require.NoError(t, err)

	ctx := context.Background()

	expected := "testValue1"
	err = repository.Set(ctx, "testKey", expected)
	require.NoError(t, err)

	actual, err := repository.Get(context.Background(), "testKey")
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func BenchmarkRedigoGet(b *testing.B) {
	repository, err := redigo.NewRepository("redis1:6379")
	require.NoError(b, err)

	ctx := context.Background()

	err = repository.Set(ctx, "testKey", "testValue")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := repository.Get(context.Background(), "testKey")
		require.NoError(b, err)
	}
}

func BenchmarkRedigoSet(b *testing.B) {
	repository, err := redigo.NewRepository("redis1:6379")
	require.NoError(b, err)

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := repository.Set(ctx, "testKey", "testValue")
		require.NoError(b, err)
	}
}

func BenchmarkRedigoGetAndSet(b *testing.B) {
	repository, err := redigo.NewRepository("redis1:6379")
	require.NoError(b, err)

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assertSetAndGet(b, repository, ctx, "testKey", "testValue")
	}
}

// -----------------------------------------------

// -------- Goredis Benchmarks -------------------

func TestGoredis(t *testing.T) {
	repository := goredis.NewRepository("redis1:6379")

	ctx := context.Background()

	expected := "testValue1"
	err := repository.Set(ctx, "testKey", expected)
	require.NoError(t, err)

	actual, err := repository.Get(context.Background(), "testKey")
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func BenchmarkGoredisGet(b *testing.B) {
	repository := goredis.NewRepository("redis1:6379")

	ctx := context.Background()

	err := repository.Set(ctx, "testKey", "testValue")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := repository.Get(context.Background(), "testKey")
		require.NoError(b, err)
	}
}

func BenchmarkGoredisSet(b *testing.B) {
	repository := goredis.NewRepository("redis1:6379")

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := repository.Set(ctx, "testKey", "testValue")
		require.NoError(b, err)
	}
}

func BenchmarkGoredisGetAndSet(b *testing.B) {
	repository := goredis.NewRepository("redis1:6379")

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assertSetAndGet(b, repository, ctx, "testKey", "testValue")
	}
}

// -----------------------------------------------

func assertSetAndGet(t testing.TB, redisRepository storage.Repository, ctx context.Context, key, value string) {
	t.Helper()

	err := redisRepository.Set(ctx, key, value)
	require.NoError(t, err)

	_, err = redisRepository.Get(ctx, key)
	require.NoError(t, err)
}
