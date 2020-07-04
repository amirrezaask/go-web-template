package http

import "github.com/labstack/echo/v4"

func withRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/", Hello())
	return e
}
