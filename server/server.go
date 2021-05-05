package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yoshikouki/semaphore-server/api"
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

	return serv.Run(conf)
}

type server struct {
	config Config
}

// Run HTTP server
func (s *server) Run(conf Config) error {
	e := echo.New()
	api.CreateEndpoints(e)

	port := fmt.Sprintf(":%d", conf.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Fatal(err)
		return err
	}
	return nil
}
