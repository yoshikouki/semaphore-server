package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/yoshikouki/semapi/model"
)

const ModelKey = "GetterModel"

func Model(model *model.Model) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			c.Set(ModelKey, model)
			return next(c)
		})
	}
}
