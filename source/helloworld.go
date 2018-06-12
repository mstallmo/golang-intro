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

	c := make(chan greeting.Salutation)
	c2 := make(chan greeting.Salutation)
	go salutations.ChannelGreeter(c)
	go salutations.ChannelGreeter(c2)

	for {
		select {
		case s, ok := <-c:
			if ok {
				fmt.Println(s, ":1")
			} else {
				return
			}
		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
			} else {
				return
			}
		default:
			fmt.Println("Waiting...")
		}
	}
}
