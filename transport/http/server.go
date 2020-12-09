package http

import (
	"app/log"

	"github.com/labstack/echo/v4"
)

func LoggerMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log.Logger().Debugf("%s\n", ctx.Request().RequestURI)
		err := h(ctx)
		if ctx.Response().Status > 499 {
			log.Logger().Errorf("%s\n", err)
		}
		log.Logger().Debugf("%s\n", err)
		return err
	}
}

func Hello() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.HTML(200, "<h1>Hello World From Golang!</h1>")
	}
}

func NewEchoServer() *echo.Echo {
	e := echo.New()
	e.Use(LoggerMiddleware)
	withRoutes(e)
	return e
}
