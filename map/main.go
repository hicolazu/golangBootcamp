package main

import "fmt"

func main() {
	var colors map[string]string
	colors2 := make(map[string]string)

	colors2["white"] = "#ffffff"

	delete(colors2, "white")

	fmt.Println(colors)
	fmt.Println(colors2)
}
