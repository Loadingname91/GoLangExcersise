package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{"http://google.com", "http://twitter.com", "http://facebook.com"}
	for _, website := range websites {
		// this will not work as the main routine would have exited before the child routine
		go checkLink(website)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Website might be down for ", link)
		return
	}
	fmt.Println("Website is okay for ", link)
}
