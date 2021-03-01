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

// *person: defines that pointerToPerson have to be a pointer of the type person
// *pointerToPerson: defines that the operation will happen with the value pointed
// by the pointer, not the adress
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
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

	// Pointer rule:
	// Turn adress to value with *adress
	// Turn value to adress with &value
	jimPointer := &jim
	jimPointer.updateName("Jimmy")

	jim.print()
}
