package goredis

import (
	"testing"
)

const (
	testKey     = "testKeyGoRedis"
	testValue   = "testValueGoRedis"
	testHashKey = "testHashGoRedis"
	testListKey = "testListGoRedis"
)

var redisRepo = NewRedisRepository("0.0.0.0:20003")

func BenchmarkGoRedisGetAndSet(b *testing.B) {
	for i := 0; i < b.N; i++ {

		// --- (1) ---
		// Set the new value into the redis
		err := redisRepo.Set(testKey, testValue)
		if err != nil {
			panic(err)
		}

		// --- (2) ---
		// Get the new value from redis
		_, err = redisRepo.Get(testKey)
		if err != nil {
			panic(err)
		}

	}
}

func BenchmarkGoRedisGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := redisRepo.Get(testKey)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGoRedisSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := redisRepo.Set(testKey, testValue)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGoRedisHashGetall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := redisRepo.HashGetAll(testHashKey)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGoRedisHashSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := redisRepo.HashSet(testHashKey, testKey, testValue)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGoRedisHashSetAndGetAllItems(b *testing.B) {

	var (
		exampleHash = map[string]string{
			"testKey1": "testValue",
			"testKey2": "testValue",
			"testKey3": "testValue",
			"testKey4": "testValue",
			"testKey5": "testValue",
		}
		hash = "BenchmarkGoRedisHashSetAndGetAllItemsHash"
	)

	for i := 0; i < b.N; i++ {
		for key, value := range exampleHash {
			err := redisRepo.HashSet(hash, key, value)
			if err != nil {
				panic(err)
			}

			_, err = redisRepo.HashGetAll(hash)
			if err != nil {
				panic(err)
			}
		}

	}
}

func BenchmarkListGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := redisRepo.ListGet(testListKey, 0, 100)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkListPushAndGet(b *testing.B) {
	slc := []string{
		"adefes",
		"fafesdfsdf",
		"dfsefefse",
		"dfsegdbfdbedb",
		"ritghtuhnert",
		"pelsfremgmuerg",
	}
	for i := 0; i < b.N; i++ {
		for _, val := range slc {
			err := redisRepo.ListPush(testListKey, val)
			if err != nil {
				panic(err)
			}
		}

		_, err := redisRepo.ListGet(testListKey, 0, len(slc))
		if err != nil {
			panic(err)
		}
	}
}
