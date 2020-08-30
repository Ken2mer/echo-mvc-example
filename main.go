package main

import (
	"fmt"
	"os"

	"github.com/Ken2mer/echo-mvc/app"
	"github.com/Ken2mer/echo-mvc/config"
)

func run() int {
	dbURL := os.Getenv("DB_URL")

	if err := app.OpenDB(dbURL); err != nil {
		fmt.Println(err)
		return 1
	}
	defer app.CloseDB()

	e := config.Routes()

	if err := app.RunServer(e); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
