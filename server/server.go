package server

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/yoshikouki/semaphore-server/api"
	"net/http"
)

func Launch() error {
	// get config
	conf, err := NewConfig()
	if err != nil {
		return err
	}

	// run Redis
	rdb, err := NewRedis(conf)
	if err != nil {
		return err
	}

	serv := &server{
		config: conf,
		redis:  rdb,
	}

	return serv.Run(conf)
}

type server struct {
	config Config
	redis  *redis.Client
}

// Run HTTP server
func (s *server) Run(conf Config) error {
	e := echo.New()
	api.DefineEndpoints(e)
	e.GET("/redis/ping", s.redisPing)

	port := fmt.Sprintf(":%d", conf.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Fatal(err)
		return err
	}
	return nil
}

func (s *server) redisPing(c echo.Context) error {
	pong, err := s.redis.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, pong)
}
