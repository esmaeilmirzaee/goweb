package main

import (
	"fmt"
	"net/http"
	"webvideos/060/controllers"
	"webvideos/060/models"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "tb"
)

func main() {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	us, err := models.NewUserService(sqlInfo)
	if err != nil {
		panic(err)
	}
	us.AutoMigrate()

	staticC := controllers.NewStatic()
	userC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")
	r.Handle("/login", userC.LoginView).Methods("GET")
	r.HandleFunc("/login", userC.Login).Methods("POST")
	fmt.Println("Listen & Serve")
	http.ListenAndServe(":3000", r)
}
