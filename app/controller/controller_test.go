package controller

import (
	"log"
	"os"
	"testing"

	"github.com/Ken2mer/echo-mvc/internal/testutils"
)

func TestMain(m *testing.M) {
	if err := testutils.SetupDB(); err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	testutils.ShutdownDB()

	os.Exit(code)
}
