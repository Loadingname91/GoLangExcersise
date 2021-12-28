package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Name   string
	Age    int
	Weight float64
}

type Car interface {
	Drive()
	Stop()
}

type lambogini struct {
	modelYear int
}

type Chevy struct {
	modelYear int
}

func PassAnything(argsValue interface{}) {
	fmt.Println("Data passed through the args Value is", argsValue)
}

func main() {

	// First way of declaring data
	var (
		m1 = 2
		m2 = 1
	)
	fmt.Println(m1+m2, "Data declared")

	// Second way of declaring data

	// var m3 int64
	// var m4 int32

	// Cannot add mismatched types
	// fmt.Println(m3 + m4)

	// Third way of declaring data

	m5 := 5
	m6 := 6

	fmt.Println(m5 + m6)

	// declaring and initializing strings

	var m7 string
	m7 = "Hello World"

	fmt.Println(m7)

	// String Methods

	// Contains

	m8 := "Data is here"
	m9 := "here"
	m10 := "XYZ"
	fmt.Println(strings.Contains(m8, m9))
	fmt.Println(strings.Contains(m8, m10))

	// Replace ALl

	m11 := "Hello Replace"
	m12 := "Replace"
	m13 := "Hitesh"

	fmt.Println(strings.ReplaceAll(m11, m12, m13))

	// Split Strings

	m14 := "Hello Guys my name is Hitesh"

	fmt.Println(strings.Split(m14, " "))

	// New Arrays

	arr := []int{1, 2, 3, 4}

	fmt.Println(arr)

	arr2 := []string{"hi", "my", "name"}

	arr2 = append(arr2, "is Hitesh")

	fmt.Println(arr2)

	// Exploring pointers

	m15 := 2
	ptr := &m15
	fmt.Println("Data present in the variable is ", *ptr, "Address is ", ptr)

	// Swaping Variables

	fmt.Println("Data present in m5 ", m5, "Data present in m6", m6)

	SwapVariables(&m5, &m6)

	fmt.Println("Data present in m5", m5, "Data present in m6", m6)

	// Declaring Struct types and working with them

	p := Person{
		Name:   "Hitesh Kumar",
		Age:    24,
		Weight: 85.6,
	}

	fmt.Println("Person data is ", p)

	fmt.Printf("Person data is %+v", p)

	// Creating a reciever for struct Person
	p.print()

	// Creating a custom interface
	lambo := lambogini{2020}
	chevy := Chevy{1998}

	lambo.Drive()
	chevy.Drive()

	// ShowCar for any car
	StartCar(lambo)
	StartCar(chevy)

	// Stop Car for any car
	StopCar(lambo)
	StopCar(chevy)

	// Passing Data through args

	PassAnything("Data good")
	PassAnything(1234)
	PassAnything(struct{}{})

	// Creating  a custom interface with make

	mymap := make(map[string]interface{})

	mymap["name"] = "Hitesh Kumar"
	mymap["value"] = "22.456"

	fmt.Println("Data", mymap)

	// Executing control statements

	f := true
	flag := &f

	if flag == nil && 2 == 3 || 3 != 5 {
		fmt.Println("Value is nil")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Data point is ", i)
	}

	// running an infinite loop
	// for {
	// 		fmt.Println("for loop is running infinitely")
	//	}

	// Parsing data through an array

	for i := range arr {
		fmt.Println("Data manipulated is ", i)
	}

	// using continue , switch , break
	flag2 := true
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			flag2 = false
			break
		} else if i == 1 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("Flag value is ", flag2)

	day := "Fri"
	switch day {
	case "Fri":
		fmt.Println("The day is Friday")
	case "Mon":
		fmt.Println("The day is Monday")
	case "Tue":
		fmt.Println("The day is Tuesday")
	case "Wed":
		fmt.Println("The day is Wednesday")
	default:
		fmt.Println("Default")
	}

	// Working with Exceptions
	// err := errors.New("The new Exception is raised")
	// switch err {
	// case err:
	// 	fmt.Println("it contains an error")
	// }

	// Running go routines 

	

}

// Another signature of the function SwapVariables (m1,m2 *int)
func SwapVariables(m1 *int, m2 *int) {
	temp := *m2
	*m2 = *m1
	*m1 = temp
}

func StartCar(c Car) {
	c.Drive()
}

func StopCar(c Car) {
	c.Stop()
}

func (p Person) print() {
	fmt.Printf("\n From the reciever person data is %+v\n", p)
}

func (l lambogini) Drive() {
	fmt.Println("Car lambogini ready to drive")
}

func (l lambogini) Stop() {
	fmt.Println("Car lambogini stopping now")
}

func (c Chevy) Drive() {
	fmt.Println("Car Chevy ready to drive")
}

func (c Chevy) Stop() {
	fmt.Println("Car chevy stopping now")
}
