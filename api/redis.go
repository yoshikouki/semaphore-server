package api

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/yoshikouki/semaphore-server/middleware"
	"net/http"
)

func redisPing(c echo.Context) error {
	rdb := c.Get(middleware.RedisKey).(*redis.Client)

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, pong)
}
