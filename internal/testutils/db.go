package testutils

import (
	"os"

	"github.com/Ken2mer/echo-mvc/app"
	"github.com/Ken2mer/echo-mvc/app/model"
)

func SetupDB() error {
	dbURL := os.Getenv("DB_URL_TEST")

	if err := app.OpenDB(dbURL); err != nil {
		return err
	}
	app.DB.AutoMigrate(&model.User{})

	return nil
}

func ShutdownDB() {
	app.DB.DropTableIfExists(&model.User{})
	app.CloseDB()
}
