package main

import "fmt"

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

type Dog struct{}

type Huskey struct {
	Speaker
}

type Speaker interface {
	Speak()
}

func (d Dog) Speak() {
	fmt.Println("Woof")
}

func main() {
	// interface related course
	h := Huskey{Cat{}}
	h.Speak()
}
