package app

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB *gorm.DB

	defaultSource = "postgres://localhost:5432/database"
)

func OpenDB(source string) error {
	var err error

	if len(source) == 0 {
		source = defaultSource
	}

	fmt.Println("Open DB")
	DB, err = gorm.Open("postgres", source)
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() {
	fmt.Println("Close DB")
	if err := DB.Close(); err != nil {
		fmt.Println(err)
	}
}
