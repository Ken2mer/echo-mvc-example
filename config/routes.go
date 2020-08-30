package config

import (
	"fmt"
	"html"
	"net/http"

	"github.com/Ken2mer/echo-mvc/app/controller"
	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Hello, %q", html.EscapeString(c.Path())))
}

func Routes(e *echo.Echo) {
	e.GET("/", hello)
	e.GET("/users", controller.Users)
	e.POST("/users", controller.CreateUser)
}
