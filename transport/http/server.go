package http

import (
	"app/monitoring"
	"github.com/labstack/echo/v4"
)


func Hello() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.HTML(200, "<h1>Hello World From Golang!</h1>")
	}
}
func LoggerMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		monitoring.Logger().Debugln(ctx.Request().RequestURI)
		err := h(ctx)
		if ctx.Response().Status > 499 {
			monitoring.Logger().Errorln(err)
		}
		monitoring.Logger().Debugln(err)
		return err
	}
}

func NewEchoServer() *echo.Echo {
	e := echo.New()
	// Add logger middleware
	e.Use(LoggerMiddleware)
	// registering routes
	e.GET("/", Hello())
	return e
}