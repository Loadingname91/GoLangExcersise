package main

func main() {
	// cards := deck{"Ace of Spades", newCard()}
	// cards = append(cards, "Ace of Hearts")
	cards := newDeck()
	cards.print("All cards in the deck")
	// fmt.Println("changing to string")
	// fmt.Println(cards.toString())
	hand, remainingDeck := deal(cards, 5)
	hand.print("cards present in the hand")
	remainingDeck.print("cards present in the remianing deck")

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
