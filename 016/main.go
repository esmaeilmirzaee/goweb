package main

import (
	"html/template"
	"os"
)

type Laptop struct {
	Count  int
	First  string
	Second string
}

type User struct {
	Name string
	Laptop
}

func main() {
	t, err := template.ParseFiles("./hello.gohtml")
	if err != nil {
		panic(err)
	}
	laptop := Laptop{
		Count:  2,
		First:  "MSI",
		Second: "Apple MacBook Pro",
	}

	data := User{
		Name:   "Esmaeil MIRZAEE",
		Laptop: laptop,
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	// Go's html template prevents code injection
	data.Name = "<script>alert(\"Hi\")</script>"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
