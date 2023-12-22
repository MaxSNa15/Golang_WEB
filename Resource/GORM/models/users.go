package models

import "GORM/data"

type User struct {
	Id       int64  `json:"id"`
	Usename  string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

func MigrarUser() {
	data.Database.AutoMigrate(User{})
}
