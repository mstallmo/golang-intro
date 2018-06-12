package main

import "./greeting"

func main() {
	mason := greeting.Name{FirstName: "Mason", LastName: "Stallmo"}
	allie := greeting.Name{FirstName: "Allie", LastName: "Heartsworm"}
	sujoy := greeting.Name{FirstName: "Sujoy", LastName: "Ganguly"}

	salutations := greeting.Salutations{
		{Name: mason, Greeting: "Hello"},
		{Name: allie, Greeting: "Hi"},
		{Name: sujoy, Greeting: "Sup"},
	}

	salutations[0].Rename("Kevin", "Stallmo")

	salutations.Greet(greeting.Println, true, 0)
}
