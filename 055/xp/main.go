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

	rows, err := db.Query("SELECT * FROM users INNER JOIN orders ON users.id=orders.user_id ORDER BY orders.user_id;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, amount int
		var name, email, desc string
		if err := rows.Scan(&id, &name, &email, &id, &id, &amount, &desc); err != nil {
			panic(err)
		}
		fmt.Println("id: ", id, "name: ", name, "email: ", email, "id: ", id, "user_id: ", id, "amount: ", amount, "desc: ", desc)
	}
	if rows.Err() != nil {
		panic(err)
	}
}
