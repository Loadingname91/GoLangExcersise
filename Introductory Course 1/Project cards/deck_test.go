package main

import (
	"os"
	"testing"
)

func TestDecks(t *testing.T) {
	deck := newDeck()
	if len(deck) != 20 {
		t.Errorf("Expected the length of cards to be 20 got : %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", deck[0])
	}

	if deck[len(deck)-1] != "Five of Clubs" {
		t.Errorf("Expected Five of Clubs but got %v", deck[len(deck)-1])
	}

}

func TestSaveFileAndReadFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveFile("_deckTesting")

	loadedDeck := readFile("_deckTesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards in the deck but got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")

}
