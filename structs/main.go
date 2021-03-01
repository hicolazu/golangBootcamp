package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}
	ana := person{firstName: "Ana", lastName: "Anderson"}

	fmt.Println(alex.firstName + " " + ana.firstName)
}
