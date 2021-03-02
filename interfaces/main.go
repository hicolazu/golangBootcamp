package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	var e englishBot
	var s spanishBot

	printGreeting(e)
	printGreeting(s)
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
