package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleFunc(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	switch r.URL.Path {
	case "/":
		fmt.Fprint(rw, "<h1>Hello Go Mux</h1>")
	case "/support":
		fmt.Fprint(rw, "<h1>Please send email to our <a href=\"support@tb.xyz\">support team</a></h1>")
	default:
		fmt.Fprint(rw, "<h3>Please check the url</h3>")
		break
	}
}

func ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("ContentType", "text/html")
	fmt.Fprint(rw, "<h1 style=\"color: red\">404</h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleFunc)
	r.HandleFunc("/support", handleFunc)
	ServeHTTP http.NotFoundHandler
	http.ListenAndServe(":3000", r)
}
