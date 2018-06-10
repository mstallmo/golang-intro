package main

import "./greeting"

func main() {
	mason := greeting.Name{FirstName: "Mason", LastName: "Stallmo"}
	s := greeting.Salutation{Name: mason, Greeting: "Hello"}
	allie := greeting.Name{FirstName: "Allie", LastName: "Heartsworm"}
	t := greeting.Salutation{Name: allie, Greeting: "Hello"}
	greeting.Greet(s, greeting.Println, true)
	greeting.Greet(t, greeting.Println, true)
}
