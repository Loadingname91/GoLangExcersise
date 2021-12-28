package main

import "fmt"

type bot interface {
	getGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func main() {

	newEnglishInstance := EnglishBot{}
	printGreeting(newEnglishInstance)

	newSpanishInstance := SpanishBot{}
	printGreeting(newSpanishInstance)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(sb SpanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (EnglishBot) getGreeting() string {
	return "Hello World"
}

func (SpanishBot) getGreeting() string {
	return "Hola"
}
