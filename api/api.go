package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func DefineEndpoints(e *echo.Echo) {
	e.GET("/ping", ping)
	e.GET("/redis/ping", redisPing)
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
