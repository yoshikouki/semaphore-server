package server

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func NewRedis(host string, port int, password string, db int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})

	return rdb, nil
}
