package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	fmt.Fprint(rw, "<h1>Welcome</h1>")
}

func contact(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	fmt.Fprint(rw, "<h3>Call us</h3>")
}

func notFound(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(http.StatusNotFound)
	fmt.Fprint(rw, "<h1>Not Found <br />404</h1>")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", router)
}
