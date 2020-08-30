package controller

import (
	"net/http"

	"github.com/Ken2mer/echo-mvc/app"
	"github.com/Ken2mer/echo-mvc/app/model"
	"github.com/labstack/echo"
)

func Users(c echo.Context) (err error) {
	var users []model.User
	app.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) (err error) {
	user := new(model.User)
	if err = c.Bind(user); err != nil {
		return
	}
	app.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}
