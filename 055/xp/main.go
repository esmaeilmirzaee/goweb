package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		}),
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	name, email := getInfo()
	user := User{
		Name:  name,
		Email: email,
	}
	if err := db.Create(&user).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", user)
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	fmt.Println("What is your email address?")
	email, _ = reader.ReadString('\n')
	return name, email
}
