package http

import "github.com/labstack/echo/v4"


type Echo struct {
    New func() *echo.Echo
}
