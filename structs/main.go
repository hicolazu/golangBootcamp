package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim.party@mail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

/*
	Value Types: Use pointers to change these things in a fuction
		- int
		- float
		- string
		- bool
		- struct

	Reference Types: Don't worry about pointers with these
		- slices
		- maps
		- channels
		- pointers
		- functions
*/
