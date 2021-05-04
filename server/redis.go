package server

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func NewRedis(conf Config) (*redis.Client, error) {
	opt := redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.RedisHost, conf.RedisPort),
		Password: conf.RedisPassword,
		DB:       conf.RedisDB,
	}

	rdb := redis.NewClient(&opt)
	return rdb, nil
}
