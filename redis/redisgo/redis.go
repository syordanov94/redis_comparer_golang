package redisgo

import (
	"rediscomparer/redis"

	redigo "github.com/gomodule/redigo/redis"
)

type RedisRepository struct {
	conn redigo.Conn
}

func NewRedisRepository(address, password string) RedisRepository {
	connection, err := redigo.Dial("tcp", address, redigo.DialPassword(password))
	if err != nil {
		panic(err)
	}
	return RedisRepository{
		conn: connection,
	}
}

var _ redis.RedisRepository = (*RedisRepository)(nil)

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
