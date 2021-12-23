package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
			cards = append(cards, values+" of "+suit)
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

func byteSliceToStringSlice(b []byte) []string {
	return strings.Split(string(b), ",")
}

func (d deck) saveFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("There was error while reading the file - ", err)
		os.Exit(1)
	}
	return byteSliceToStringSlice(bs)
}

func (d deck) shuffle() deck {
	// creating a random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
