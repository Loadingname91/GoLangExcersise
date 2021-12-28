package main

import "fmt"

type ContactInfo struct {
	emailId string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func main() {
	// another way
	alex1 := Person{"Alex", "Anderson", ContactInfo{"AlexAnderson@gamil.com", 560072}}
	fmt.Println("first way to print", alex1)
	alex2 := Person{firstName: "Alex", lastName: "Anderson", contact: ContactInfo{"alexAnderson@yahoo.com", 560073}}
	fmt.Printf("Second way to represent the contents %+v\n", alex2)

	// initial value if not initialized
	var alex3 Person
	fmt.Printf("if the data is not initialized %+v", alex3)

	fmt.Println("Declaring contact information")

	sherly := Person{
		firstName: "Sherly",
		lastName:  "Smith",
		contact: ContactInfo{
			emailId: "sheryly@gmail.com",
			zipCode: 560040,
		},
	}

	// 1st way
	// sherlyPointer := &sherly

	// sherlyPointer.passbyreferenceUpdateName("Shrely", "anderson")

	// 2nd way
	sherly.passbyreferenceUpdateName("Shrely", "Anderson")
	sherly.print()

	// an exception to pass by value is when we are passing in values to a function of the type slice

	newSlice := []int{0, 1, 2, 4, 5}
	updateSlice(newSlice)
	fmt.Println("data present in the updated slice is ", newSlice)

}
func updateSlice(s []int) {
	s[0] = 11
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

func (p Person) passByValueUpdateName(newFirstName string, newLastName string) {
	p.firstName = newFirstName
	p.lastName = newLastName

}

func (pointerToPerson *Person) passbyreferenceUpdateName(newFirstName string, newLastName string) {
	(*pointerToPerson).firstName = newFirstName
	(*pointerToPerson).lastName = newLastName
}
