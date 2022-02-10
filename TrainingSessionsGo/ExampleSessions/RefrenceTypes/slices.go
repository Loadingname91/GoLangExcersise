package main

import "fmt"

func main() {

	sample := make([]int, 3, 5)
	sample2 := append(sample, 40, 50)
	sample3 := append(sample, 40, 50, 60, 60, 60)
	fmt.Println(sample2, sample, sample3)
	fmt.Println(cap(sample2), cap(sample), cap(sample3))

	foo := []int{1, 2, 3} // This is a slice
	foo = append(foo, 4, 5, 6)
	fmt.Println(foo)

	sample4 := make([]int, 3, 5)
	// The below line breaks the code
	// sample4[4] = 10
	fmt.Println(sample4)

	a := make([]int, 3, 5)
	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Println("Capacity is ", cap(a), "length is", len(a), a)
	a = append(a, 4, 5, 6)
	fmt.Println("Capacity is ", cap(a), "length is", len(a), a)

	a = append(a, 7, 8)
	fmt.Println("Capacity is ", cap(a), "length is", len(a), a)

	a = append(a, 9, 10)
	fmt.Println("Capacity is ", cap(a), "length is", len(a), a)

	b := []int{0: 1, 1: 2}
	fmt.Println(b)

}
