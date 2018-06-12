package main

import "./greeting"

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

	RenameToFrog(&salutations[0])

	salutations.Greet(greeting.Println, false, 0)
}
