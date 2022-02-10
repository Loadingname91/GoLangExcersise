package main

import "fmt"

type Car struct {
	Model string
}

func main() {
	// unbuffered channels
	// c := make(chan int)
	// go func() {
	// 	c <- 1
	// }()

	// val := <-c
	// fmt.Println(val)
	// go func() {
	// 	c <- 2
	// }()
	// fmt.Println(2)

	// buffered channels

	c := make(chan int, 3)

	go func() {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}

	c2 := make(chan *Car, 3)

	go func() {
		c2 <- &Car{Model: "1"}
		c2 <- &Car{Model: "2"}
		c2 <- &Car{Model: "3"}
		c2 <- &Car{Model: "4"}
		close(c2)
	}()

	for i := range c2 {
		fmt.Println(i.Model)
	}

}
