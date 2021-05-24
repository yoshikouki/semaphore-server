package server

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/yoshikouki/semaphore-server/api"
	"github.com/yoshikouki/semaphore-server/middleware"
)

func Launch(cfg Config) error {
	// get config
	conf, err := NewConfig(cfg)
	if err != nil {
		return err
	}

	// run Redis
	rdb, err := NewRedis(
		conf.RedisHost,
		conf.RedisPort,
		conf.RedisPassword,
		conf.RedisDB,
	)
	if err != nil {
		return err
	}

	return serverRun(conf, rdb)
}

// Run HTTP server
func serverRun(conf Config, redis *redis.Client) error {
	e := echo.New()
	e.Use(middleware.Redis(redis))

	api.DefineEndpoints(e)

	port := fmt.Sprintf(":%d", conf.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Fatal(err)
		return err
	}
	return nil
}
