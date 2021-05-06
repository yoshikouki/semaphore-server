package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type response struct {
	message string
}

// lockIfNotExists is Mutex what can only be used to maintain atomicity, if key don't exists.
// key: `org-repo-stageName`
func lockIfNotExists(c echo.Context) error {
	return c.String(http.StatusOK, "test: #lockIfNotExists")
}
