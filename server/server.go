package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Launch() error {
	// get config
	conf, err := NewConfig()
	if err != nil {
		return err
	}

	serv := &server{
		config: conf,
	}

	return serv.Run()
}

type server struct {
	config Config
}

// Run HTTP server
func (s *server) Run() error {
	e := echo.New()
	e.GET("/ping", ping)

	port:= fmt.Sprintf(":%d", s.config.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Fatal(err)
		return err
	}
	return nil
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
