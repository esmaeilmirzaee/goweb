package main

import (
	"fmt"
	"goweb/060/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "tb"
)

func main() {
	// Checking remember token
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	us, err := models.NewUserService(sqlInfo)
	if err != nil {
		panic(err)
	}
	us.DestructiveReset()

	user := models.User{
		Name:     "E E",
		Email:    "e@e.e",
		Password: "1234",
		Remember: "abc123",
	}
	err = us.Create(&user)
	fmt.Printf("%+v", user)
	user2, err := us.ByRemember("abc123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", *user2)
}
