package main

import (
	"fmt"
	"webvideos/060/rand"
)

func main() {
	fmt.Println(rand.String(10))
	fmt.Println(rand.RememberToken())
}
