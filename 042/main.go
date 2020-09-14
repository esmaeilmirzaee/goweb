package main

import (
	"fmt"
	"net/http"
	"webvideos/042/controllers"

	"github.com/gorilla/mux"
)

func main() {
	usersC := controllers.NewUsers()
	staticC := controllers.NewStatic()
	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	fmt.Println("Listen & Serve")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
