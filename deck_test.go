package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	testDeck := newDeck()

	if len(testDeck) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(testDeck))
	}

	if testDeck[0] != "Two of Spades" {
		t.Errorf("Expected the first card to be Two of Spades but got %v", testDeck[0])
	}

	if testDeck[len(testDeck)-1] != "Ace of Clubs" {
		t.Errorf("Expected the last card to be Ace of Clubs got got %v", testDeck[len(testDeck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
