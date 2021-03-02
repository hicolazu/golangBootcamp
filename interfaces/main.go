package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	var e englishBot
	var s spanishBot

	e.printGreeting()
	s.printGreeting()
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (e englishBot) printGreeting() {
	fmt.Println(e.getGreeting())
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (s spanishBot) printGreeting() {
	fmt.Println(s.getGreeting())
}
