package redigo

import (
	"context"

	"rediscomparer/storage"

	redigo "github.com/gomodule/redigo/redis"
)

type Repository struct {
	conn redigo.Conn
}

var _ storage.Repository = (*Repository)(nil)

func NewRepository(address string) (*Repository, error) {
	connection, err := redigo.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Repository{
		conn: connection,
	}, nil
}

func (r *Repository) Get(ctx context.Context, key string) (string, error) {
	return redigo.String(r.conn.Do("GET", key))
}

func (r *Repository) Set(ctx context.Context, key, value string) error {
	_, err := r.conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}
