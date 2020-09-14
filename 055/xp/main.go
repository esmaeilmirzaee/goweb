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

	var id int
	err = db.QueryRow(`INSERT INTO users(name, email) VALUES($1, $2) RETURNING id`, "E E", "e@e.e").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
