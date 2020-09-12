package main

import (
	"fmt"
	"net/http"
	"webvideos/037/views"

	"github.com/gorilla/mux"
)

var (
	homeView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	fmt.Println("Listen & Serve")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
