package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func New(e *echo.Echo) {
	e.GET("/ping", ping)
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
