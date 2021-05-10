package model

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Model struct {
	redis *redis.Client
}

func NewModel(rdb *redis.Client) (*Model, error) {
	model := &Model{
		redis: rdb,
	}

	ctx := context.Background()
	_, err := model.redis.Ping(ctx).Result()

	return model, err
}
