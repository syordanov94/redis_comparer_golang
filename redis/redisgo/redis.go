package redisgo

import (
	redigo "github.com/gomodule/redigo/redis"
)

type RedisRepository struct {
	conn redigo.Conn
}

func NewRedisRepository(address string) RedisRepository {
	connection, err := redigo.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	return RedisRepository{
		conn: connection,
	}
}

func (repo *RedisRepository) Get(key string) (string, error) {
	val, err := redigo.String(repo.conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return val, nil
}

func (repo *RedisRepository) Set(key, value string) error {
	_, err := repo.conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RedisRepository) HashGetAll(key string) (map[string]string, error) {
	val, err := redigo.StringMap(repo.conn.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (repo *RedisRepository) HashSet(hashkey, key, value string) error {
	_, err := repo.conn.Do("HSET", hashkey, key, value)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RedisRepository) ListGet(key string, start, end int) ([]string, error) {
	val, err := redigo.Strings(repo.conn.Do("LRANGE", key, start, end))
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (repo *RedisRepository) ListPush(key, value string) error {
	_, err := repo.conn.Do("LPUSH", key, value)
	if err != nil {
		return err
	}
	return nil
}
