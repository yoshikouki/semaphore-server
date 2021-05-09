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
	LockTarget string        `form:"lock_target" json:"lock_target"`
	User       string        `form:"user" json:"user"`
	TTL        time.Duration `form:"ttl" json:"ttl"`
}

// lockIfNotExists is Mutex what can only be used to maintain atomicity, if key don't exists.
// key: `org-repo-stage`
func lockIfNotExists(c echo.Context) error {
	params := LockIfNotExistsParams{}

	err := echo.FormFieldBinder(c).
		MustString("lock_target", &params.LockTarget).
		MustString("user", &params.User).
		MustDuration("ttl", &params.TTL).
		BindErrors()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, params)
}
