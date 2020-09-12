package main

import (
	"fmt"
	"net/http"
	"webvideos/042/controllers"
	"webvideos/042/views"

	"github.com/gorilla/mux"
)

var (
	homeView   *views.View
	signupView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	usersC := controllers.NewUsers()
	r := mux.NewRouter()
	r.HandleFunc("/", home)
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
