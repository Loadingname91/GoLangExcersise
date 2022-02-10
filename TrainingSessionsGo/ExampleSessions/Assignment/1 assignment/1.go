package main

import (
	"errors"
	"fmt"
)

func addItem(k, v string) error {
	// Add data to the map
	if _, ok := data[k]; ok {
		// fmt.Println(k, " Key already exists")
		return errors.New("Key already exists")
	}
	data[k] = v
	// fmt.Println(k, "Key added with value", v)
	return nil
}

func updateItem(k, v string) {
	// update data to the map
	if _, ok := data[k]; ok {
		fmt.Println(k, " Key does not exists")
	}
	data[k] = v
	result := fmt.Sprintf("%s Key updated with value %s", k, v)
	fmt.Println(result)

}

func getById(k string) {
	// Get data from the map
	if res, ok := data[k]; ok {
		result := fmt.Sprintf("key %s has the value %s", k, res)
		fmt.Println(result)
	} else {
		fmt.Println("Key does not exists")
	}
}

func getAll() {
	// Get all data from the map
	if len(data) == 0 {
		fmt.Println("Map is empty")
		return
	}
	fmt.Println("Reterving all data from the map")
	for k, v := range data {
		fmt.Println("key :", k, "value :", v)
	}
}

func deleteItem(k string) {
	// delete item from the map
	if _, ok := data[k]; ok {
		fmt.Println("Key does not exists")
		return
	}
	delete(data, k)
}

// Creating a map
var data map[string]string

// init function is used to declare the data before the main function is called.
func init() {
	data = make(map[string]string)
}

func main() {
	// fmt.Println("inital data in the map", data)

	// Add data to the map
	_ = addItem("5", "Hello")
	// if err != nil {
	// 	fmt.Println("Error occured", err)
	// 	return
	// }
	_ = addItem("5", "World")

	addItem("5", "Hello")
	// if err1 != nil {
	// 	fmt.Println("Error occured", err)
	// 	return
	// }
	// Get all data from the map
	getAll()

	// Update item to the map
	updateItem("10", "World2") // this will break the code

	// Get data from the map
	updateItem("5", "Hello2")

	// Get all data from the map
	getAll()

	getById("5")
	getById("6")

	// Get all data from the map
	getAll()

	// Delete item from the map
	deleteItem("5")

	deleteItem("11")

	// Get all data from the map
	getAll()
}
