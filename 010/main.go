package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(rw, "Hello Go")
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	http.ListenAndServe(":3000", router)
}
