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

type Renameable interface {
	Rename(newName Name)
}

func (salutation *Salutation) Rename(newName Name) {
	salutation.Name = newName
}

func (salutation *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.Rename(Name{FirstName: s, LastName: s})
	n = len(s)
	err = nil
	return
}

type Salutations []Salutation

type Printer func(string)

func (salutations Salutations) Greet(do Printer, isFormal bool, times int) {
	for _, element := range salutations {
		if prefix := getPrefix(element.Name.FirstName); isFormal {
			element.Name.LastName = prefix + element.Name.LastName
			message, _ := CreateMessage(element.Name, element.Greeting)
			do(message)
		} else {
			_, alternate := CreateMessage(element.Name, element.Greeting)
			do(alternate)
		}
	}
}

func (salutations Salutations) ChannelGreeter(c chan Salutation) {
	for _, s := range salutations {
		c <- s
	}
	close(c)
}

func getPrefix(name string) (prefix string) {
	prefixMap := map[string]string{
		"Mason":  "Mr. ",
		"Sujoy":  "Dr. ",
		"Allie":  "Ms. ",
		"Nathan": "Dr. ",
	}

	return prefixMap[name]
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
