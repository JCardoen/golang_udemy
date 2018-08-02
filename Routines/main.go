package main

import (
	"net/http"
	"fmt"
	"time"
)

func main() {
	links := []string {
		"https://google.com",
		"https://diffdigital.be",
		"https://golang.org",
		"https://facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go check_link(link, c)
	}

	// Infinite loop
	// Channels: data collection from go routines
	for l := range c {
		// Function literal to provide instant access to the links
		go func(link string) {
			// Sleep for 5 seconds
			time.Sleep(5 * time.Second)
			check_link(link, c)
		}(l)
	}
}

func check_link(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
	} else {
		fmt.Println(link, " is up")
		c <- link
	}
}