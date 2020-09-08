package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rw.Header().Set("Content-Type", "text/html")
	fmt.Fprint(rw, "<h1>Welcome</h1>")
}

func notFound(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(http.StatusNotFound)
	fmt.Fprint(rw, "<h3>Not Found</h3>")
}

func main() {
	router := httprouter.New()
	router.GET("/", home)
	router.NotFound = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", router)
}
