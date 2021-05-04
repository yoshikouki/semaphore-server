package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Launch() error {
	e := echo.New()
	e.GET("/ping", pong)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Fatal(err)
		return err
	}
	return nil
}

func pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
