package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	/* for {
		go checkLink(<-c, c)
	} */

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

// go -> routine
//
//							 CPU Core
// 						 Go Scheduler
// Go Routine - Go Routine - Go Routine
//
// Concurrency: if some go routine has a blocking code - such as a http request -
// the Go Scheduler will execute the other routine
//
// chan -> Channel
// Channels: way of comunnication between the go routines and the main go routine
