package main

import (
	"fmt"
	"sync"
)

func main() {

	/*

		Lambda function
		func(){
			fmt.Println("Hello World")
		}()
		func(){
			fmt.Println("Hello World")
		}()

	*/

	// wait groups
	wg := sync.WaitGroup{}
	// Use of mutexes to make sure to carefull access critical section
	wg.Add(2)
	go func() {
		fmt.Println("Hello World")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello World")
		wg.Done()
	}()
	fmt.Println("Done")
	wg.Wait()
	fmt.Println("Fin")
}
