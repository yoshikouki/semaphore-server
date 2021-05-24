package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/yoshikouki/semaphore-server/middleware"
	"github.com/yoshikouki/semaphore-server/model"
	"net/http"
	"strconv"
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
	GetLocked  string `json:"getLocked"`
	User       string `json:"user"`
	ExpireDate string `json:"expireDate"`
}

// lockIfNotExists is Mutex what can only be used to maintain atomicity, if key don't exists.
// key: `org-repo-stage`
func lockIfNotExists(c echo.Context) error {
	params := LockIfNotExistsParams{}
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
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	getLocked, user, expireDate, err := m.LockIfNotExists(ctx, params.LockTarget, params.User, ttl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := map[string]string{
		"getLocked":  strconv.FormatBool(getLocked),
		"user":       user,
		"expireDate": expireDate.Format("2006/01/02 15:04:05"),
	}

	return c.JSON(http.StatusOK, res)
}

type unlockParams struct {
	UnlockTarget string `json:"unlock_target" validate:"required"`
	User         string `json:"user" validate:"required"`
}

type UnlockResponse struct {
	GetUnlock string `json:"getUnlock"`
	Message   string `json:"message"`
}

func unlock(c echo.Context) error {
	params := &unlockParams{}
	m := c.Get(middleware.ModelKey).(*model.Model)
	ctx := context.Background()

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	getUnlock, message, err := m.Unlock(ctx, params.UnlockTarget, params.User)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res := map[string]string{
		"getUnlock": strconv.FormatBool(getUnlock),
		"message":   message,
	}

	return c.JSON(http.StatusOK, res)
}
