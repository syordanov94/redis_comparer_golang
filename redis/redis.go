package redis

type RedisRepository interface {
	Get(key string) (string, error)
	Set(key, value string) error
}
