package main

import (
	"fmt"

	"./greeting"
)

func RenameToFrog(r greeting.Renameable) {
	r.Rename(greeting.Name{FirstName: "Frog", LastName: "McFroggerson"})
}

func main() {
	mason := greeting.Name{FirstName: "Mason", LastName: "Stallmo"}
	allie := greeting.Name{FirstName: "Allie", LastName: "Heartsworm"}
	sujoy := greeting.Name{FirstName: "Sujoy", LastName: "Ganguly"}

	salutations := greeting.Salutations{
		{Name: mason, Greeting: "Hello"},
		{Name: allie, Greeting: "Hi"},
		{Name: sujoy, Greeting: "Sup"},
	}

	fmt.Fprintf(&salutations[0], "The count is %d", 10)

	done := make(chan bool)

	go func() {
		salutations.Greet(greeting.CreatePrintFunction(" <C>"), false, 0)
		done <- true
	}()
	salutations.Greet(greeting.Println, false, 0)
	<-done
}
