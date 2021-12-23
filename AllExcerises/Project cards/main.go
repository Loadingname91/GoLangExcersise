package main

import "fmt"

func main() {
	// cards := deck{"Ace of Spades", newCard()}
	// cards = append(cards, "Ace of Hearts")
	cards := newDeck()
	cards.print("All cards in the deck")
	// Saving deck to the store
	cards.saveFile("MyDeck")
	// fmt.Println("changing to string")
	// fmt.Println(cards.toString())
	hand, remainingDeck := deal(cards, 5)
	hand.print("cards present in the hand")
	remainingDeck.print("cards present in the remianing deck")

	fmt.Println("\nloading Deck from Memory\n")
	loadedDeck := readFile("MyDeck")
	// bad file loading example
	// loadedDeck := readFile("NonExistent")
	loadedDeck.print("\nData loaded from the deck\n")

	// shuffle deck

	shuffled := cards.shuffle()
	shuffled.print("Shuffled cards present in the deck")

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
