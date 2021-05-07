package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type response struct {
	message string
}

type LockIfNotExistsParams struct {
	LockTarget string `json:"lock_target"`
	User       string `json:"user"`
	TTL        string `json:"ttl"`
}

// lockIfNotExists is Mutex what can only be used to maintain atomicity, if key don't exists.
// key: `org-repo-stage`
func lockIfNotExists(c echo.Context) error {
	params := LockIfNotExistsParams{}
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	t, err := time.ParseDuration(params.TTL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, t.String())
}
