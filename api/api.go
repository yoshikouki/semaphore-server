package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func DefineEndpoints(e *echo.Echo) {
	e.GET("/semapi/ping", ping)
	e.GET("/semapi/redis/ping", redisPing)
	e.POST("/semapi/lock", lockIfNotExists)
	e.POST("/semapi/unlock", unlock)
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
