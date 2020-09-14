package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "password"
	dbname   = "tb"
)

func main() {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 1; i < 7; i++ {
		userId := 0
		if i > 3 {
			userId = 6
		}
		amount := 100 * i
		desc := fmt.Sprintf("USB-C Adapter x%d", i)
		_, err := db.Exec(`INSERT INTO orders(user_id, amount, description) VALUES($1, $2, $3)`, userId, amount, desc)
		if err != nil {
			panic(err)
		}
	}
}
