package server

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Launch() error {
	// set config
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
		redis: rdb,
	}

	return serv.Run()
}

type server struct {
	config Config
	redis *redis.Client
}

// Run HTTP server
func (s *server) Run() error {
	e := echo.New()
	e.GET("/ping", pong)

	port:= fmt.Sprintf(":%d", s.config.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Fatal(err)
		return err
	}
	return nil
}

func pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
