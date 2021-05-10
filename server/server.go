package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yoshikouki/semaphore-server/api"
	"github.com/yoshikouki/semaphore-server/middleware"
	"github.com/yoshikouki/semaphore-server/model"
)

func Launch() error {
	// get config
	conf, err := NewConfig()
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

	model, err := model.NewModel(rdb)
	if err != nil {
		return err
	}

	return serverRun(conf, model)
}

// Run HTTP server
func serverRun(conf Config, model *model.Model) error {
	e := echo.New()
	e.Use(middleware.Model(model))
	e.Validator = middleware.NewCustomValidator()

	api.DefineEndpoints(e)

	port := fmt.Sprintf(":%d", conf.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Fatal(err)
		return err
	}
	return nil
}
