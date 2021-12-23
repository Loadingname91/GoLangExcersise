package main

import (
	"fmt"
	"strings"
)

// create a type called deck
// which would house a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, suit+"of"+values)
		}
	}

	return cards

}

func (de deck) print(description string) {
	fmt.Println(description)
	for i, card := range de {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}
