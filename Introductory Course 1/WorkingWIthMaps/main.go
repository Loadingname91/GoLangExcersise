package main

import "fmt"

func main() {
	// 1st way
	// colors := map[string]string{
	// 	"red":    "Red Color",
	// 	"yellow": "Yellow Color",
	// }
	// 2nd way
	// var colors map[string]string
	// 3rd way
	colors := make(map[string]string)

	colors["white"] = "White Color"

	fmt.Println(colors)

	fmt.Println("Color value inside map `color` is ", colors["white"])

	// Deleting keys inside the map

	delete(colors, "white")

	fmt.Println("keys present in the map are ", colors)

	// iterating over maps

	newColors := map[string]string{
		"Red":    "Red Color",
		"White":  "White Color",
		"Orange": "Orange Color",
		"Pink":   "Pink Color",
	}

	printDataInsideMaps(newColors)
	fmt.Println("Value of White is updated ", newColors)
}

func printDataInsideMaps(data map[string]string) {
	for key, value := range data {
		if key == "White" {
			data["White"] = "Change color to White"
		}
		fmt.Printf("data present in the map is `%+v` with its key being `%+v` \n", value, key)
	}
}
