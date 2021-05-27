package api

import (
	"github.com/labstack/echo/v4"
	"github.com/yoshikouki/semapi/middleware"
	"github.com/yoshikouki/semapi/model"
	"net/http"
)

func redisPing(c echo.Context) error {
	m := c.Get(middleware.ModelKey).(*model.Model)

	pong, err := m.RedisPing()
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, pong)
}
