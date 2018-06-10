package main

import "./greeting"

func main() {
	var s = greeting.Salutation{"Mason", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("?"), true)
}
