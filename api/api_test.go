package api

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func dummyContext(
	t *testing.T,
	reqType string,
	reqPath string,
	args interface{},
) (echo.Context, *httptest.ResponseRecorder) {
	var rp []byte
	var jsonErr error

	if args != nil {
		switch args.(type) {
		case string:
			rp = []byte(args.(string))
		default:
			rp, jsonErr = json.Marshal(args)
			assert.NoError(t, jsonErr)
		}
	}

	req := httptest.NewRequest(reqType, reqPath, bytes.NewReader(rp))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)

	// WIP: Create Stub
	//b, _ := model.NewModelDummy()
	//ctx.Set(middleware.ModelKey, b)

	return ctx, rec
}

func Test_ping(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/semaphore/ping", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, ping(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "pong", rec.Body.String())
	}
}
