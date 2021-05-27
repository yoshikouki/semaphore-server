package model

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Model struct {
	redis     *redis.Client
	Semaphore Semaphore
}

func NewModel(rdb *redis.Client) (*Model, error) {
	model := &Model{
		redis:     rdb,
		Semaphore: Semaphore{},
	}

	ctx := context.Background()
	_, err := model.redis.Ping(ctx).Result()

	return model, err
}

func (m *Model) RedisPing(ctx context.Context) (string, error) {
	return m.redis.Ping(ctx).Result()
}
