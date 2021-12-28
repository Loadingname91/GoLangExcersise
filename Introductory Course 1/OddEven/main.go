package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, val := range numbers {
		if val%2 == 0 {
			fmt.Println(strconv.Itoa(val) + " is " + "even")
		} else {
			fmt.Println(strconv.Itoa(val) + " is " + "odd")
		}
	}
}
