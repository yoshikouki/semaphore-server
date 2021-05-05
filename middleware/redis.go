package middleware

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

const RedisKey = "GetterRedis"

func Redis(rdb *redis.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			c.Set(RedisKey, rdb)
			return next(c)
		})
	}
}
