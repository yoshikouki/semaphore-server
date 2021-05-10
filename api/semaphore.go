package api

import (
	"github.com/labstack/echo/v4"
	"github.com/yoshikouki/semaphore-server/middleware"
	"github.com/yoshikouki/semaphore-server/model"
	"net/http"
	"time"
)

type response struct {
	message string
}

type LockIfNotExistsParams struct {
	LockTarget string `json:"lock_target" validate:"required"`
	User       string `json:"user" validate:"required"`
	TTL        string `json:"ttl" validate:"required"`
}

type LockIfNotExistsResponse struct {
	isLocked   bool
	user       string
	expireDate time.Time
}

// lockIfNotExists is Mutex what can only be used to maintain atomicity, if key don't exists.
// key: `org-repo-stage`
func lockIfNotExists(c echo.Context) error {
	params := LockIfNotExistsParams{}
	m := c.Get(middleware.ModelKey).(*model.Model)

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	t, err := time.ParseDuration(params.TTL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	isLocked, user, expireDate, err := m.Semaphore.SetIfNotExists(params.LockTarget, params.User, t)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := LockIfNotExistsResponse{
		isLocked:   isLocked,
		user:       user,
		expireDate: expireDate,
	}

	return c.JSON(http.StatusOK, res)
}
