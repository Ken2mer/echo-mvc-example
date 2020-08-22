package model

type User struct {
	UserID uint   `json:"user_id" gorm:"primary_key"`
	Name   string `json:"name"`
}
