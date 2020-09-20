package main

import (
	"fmt"
	"goweb/060/models"

	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "password"
	dbname   = "tb"
)

type User struct {
	gorm.Model
	Email string `gorm:"not null;unique"`
	Name  string
}

func main() {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	us, err := models.NewUserService(sqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()

	user, err := us.ByID(2)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	// if err := us.DestructiveReset(false); err != nil {
	// panic(err)
	// }
}
