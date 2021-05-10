package api

import (
	"github.com/labstack/echo/v4"
	"github.com/yoshikouki/semaphore-server/middleware"
	"github.com/yoshikouki/semaphore-server/model"
	"net/http"
)

func redisPing(c echo.Context) error {
	m := c.Get(middleware.ModelKey).(model.Model)

	pong, err := m.RedisPing()
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, pong)
}
