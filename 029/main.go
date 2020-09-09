package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var contactTemplate *template.Template

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	contactTemplate, err = template.ParseFiles("029/contact.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/contact", contact)
	fmt.Println("Listen and serve on port 3000")
	http.ListenAndServe(":3000", r)
}
