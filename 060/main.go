package main

import (
	"fmt"
	"net/http"
	"webvideos/060/controllers"

	"github.com/gorilla/mux"
)

func main() {
	userC := controllers.NewUsers()
	staticC := controllers.NewStatic()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")
	fmt.Println("Listen & Serve")
	http.ListenAndServe(":3000", r)
}
