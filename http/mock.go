package http

import (
	"github.com/labstack/echo/v4"
	"github.com/victorhsb/go-mocker/model"
)

func Handler(route *model.Route) func(c echo.Context) error {
	return func(c echo.Context) error {
		if route == nil {
			return c.String(404, "route not found")
		}
		return c.String(route.Output.StatusCode, route.Output.Body.String())
	}
}
