package config

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Middlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
}
