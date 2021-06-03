package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/yoshikouki/semapi/middleware"
	"github.com/yoshikouki/semapi/model"
	"net/http"
	"time"
)

type response struct {
	message string
}

type LockParams struct {
	Target string `json:"lock_target" validate:"required"`
	User   string `json:"user" validate:"required"`
	TTL    string `json:"ttl" validate:"required"`
}

// lock is Mutex what can only be used to maintain atomicity, if key don't exists.
// key: `org-repo-stage`
func lock(c echo.Context) error {
	params := LockParams{}
	m := c.Get(middleware.ModelKey).(*model.Model)
	ctx := context.Background()

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ttl, err := time.ParseDuration(params.TTL)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = m.Lock(ctx, params.Target, params.User, ttl)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "OK")
}

type UnlockParams struct {
	Target string `json:"unlock_target" validate:"required"`
	User   string `json:"user" validate:"required"`
}

func unlock(c echo.Context) error {
	params := &UnlockParams{}
	m := c.Get(middleware.ModelKey).(*model.Model)
	ctx := context.Background()

	if err := c.Bind(&params); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(params); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err := m.Unlock(ctx, params.Target, params.User)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "OK")
}
