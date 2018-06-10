package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

const (
	PI       = 3.14
	Language = "Go"
)

const (
	A = iota
	B
	C
)

func main() {
	var s = Salutation{}
	s.name = "Mason"
	s.greeting = "Hello!"

	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
