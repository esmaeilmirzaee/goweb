package main

import (
	"fmt"
	"net/http"
	"webvideos/035/views"

	"github.com/gorilla/mux"
)

var (
	homeView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	fmt.Println("Listen & Serve")
	http.ListenAndServe(":3000", r)
}
