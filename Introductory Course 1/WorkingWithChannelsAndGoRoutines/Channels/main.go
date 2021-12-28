package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{"http://google.com", "http://twitter.com", "http://facebook.com"}
	channelInstance := make(chan string)
	for _, website := range websites {
		// this will not work as the main routine would have exited before the child routine
		go checkLink(website, channelInstance)
	}
	// for i := 0; i < len(websites); i++ {
	// 	// fmt.Println(<-channelInstance)

	// }
	for l := range channelInstance {
		go func(l string) {
			time.Sleep(time.Second * 5)
			checkLink(l, channelInstance)
		}(l)
	}
}

func checkLink(link string, channelInstance chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Website might be down for ", link)
		// channelInstance <- "The Webiste might be down"
		channelInstance <- link
		return
	}
	fmt.Println("Website is okay for ", link)
	// channelInstance <- "The website is all okay"
	channelInstance <- link
}
