package greeting

import "fmt"

type Salutation struct {
	Name     Name
	Greeting string
}

type Name struct {
	FirstName string
	LastName  string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	if prefix := getPrefix(salutation.Name.FirstName); isFormal {
		salutation.Name.LastName = prefix + salutation.Name.LastName
		message, _ := CreateMessage(salutation.Name, salutation.Greeting)
		do(message)
	} else {
		_, alternate := CreateMessage(salutation.Name, salutation.Greeting)
		do(alternate)
	}
}

func getPrefix(name string) (prefix string) {
	switch name {
	case "Mason":
		prefix = "Mr. "
	case "Sujoy":
		prefix = "Dr. "
	case "Allie":
		prefix = "Ms. "
	default:
		prefix = "Dude "
	}

	return
}

func CreateMessage(name Name, greeting string) (message string, alternate string) {
	message = greeting + " " + name.LastName
	alternate = "HEY!" + " " + name.FirstName
	return
}

func Print(s string) {
	fmt.Print(s)
}

func Println(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}