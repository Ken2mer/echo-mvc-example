package app

import (
	"github.com/labstack/echo"
)

func RunServer(e *echo.Echo) error {
	if err := e.Start(":1323"); err != nil {
		return err
	}
	return nil
}
