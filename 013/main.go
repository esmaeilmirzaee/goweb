package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rw.Header().Set("ContentType", "text/html")
	fmt.Fprint(rw, "<h1 style=\"color: red\">Welcome</h1>")
}

func contact(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rw.Header().Set("ContentType", "text/html")
	fmt.Fprint(rw, "<h3>Contact</h3>")
}

func main() {
	router := httprouter.New()
	router.GET("/", home)
	router.GET("/contact", contact)
	http.ListenAndServe(":3000", router)
}
