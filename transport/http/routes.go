package http

import "github.com/labstack/echo/v4"


func Hello() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.HTML(200, "<h1>Hello World From Golang!</h1>")
	}
}
