package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/yoshikouki/semapi/middleware"
	"github.com/yoshikouki/semapi/model"
	"net/http"
)

func DefineEndpoints(e *echo.Echo) {
	e.GET("/semapi/health-check", healthCheck)
	e.POST("/semapi/lock", lockIfNotExists)
	e.POST("/semapi/unlock", unlock)
}

func healthCheck(c echo.Context) error {
	ctx := context.Background()
	m := c.Get(middleware.ModelKey).(*model.Model)

	_, err := m.RedisPing(ctx)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "pong")
}
