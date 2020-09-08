package main

import (
	"html/template"
	"os"
)

type User struct {
	Name    string
	Classes []string
}

func main() {
	t, err := template.ParseFiles("./temp.gohtml")
	if err != nil {
		panic(err)
	}
	classes := []string{"a", "b"}
	data := User{
		Name:    "Esmaeil MIRZAEE",
		Classes: classes,
	}
	t.Execute(os.Stdout, data)
}
