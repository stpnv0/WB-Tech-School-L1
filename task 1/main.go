package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

func (h *Human) PrintAge() {
	fmt.Println(h.Age)
}

func (h *Human) ChangeName(newName string) {
	h.Name = newName
}

type Action struct {
	Human
}

func main() {
	person := Action{
		Human{
			Name:    "Dmitry",
			Surname: "Stepanov",
			Age:     18,
		},
	}
	person.PrintAge()
	person.ChangeName("DimAss")
	fmt.Println(person.Name)
}
