package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of "deck"
// And this deck is a slice of strings
type deck []string

// This function makes a new slice of strings that is of the type 'deck'.
func newDeck() deck {
	cards := deck{}

	cardSuits := deck{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := deck{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Prints all of the cards in a given deck
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

// Deals the cards using the corresponding deck and the given hand size
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Converts a slice of strings to single string that contains every string of the slice.
func (d deck) toString() string {

	// Below is another way of doing the type conversion. I'm not sure is type conversion necessary because the type deck is already a string
	// return strings.Join([]string(d), ",")

	return strings.Join(d, ",")
}

// Saves a deck with a given filename to a separate file.
func (d deck) saveToFile(filename string) error {
	deckToString := d.toString()
	deckStringToByte := []byte(deckToString)
	return ioutil.WriteFile(filename, deckStringToByte, 0666)
}

// Creates a new deck from a file that has been created by using the 'saveToFile' method.
func newDeckFromFile(filename string) deck {
	byteSlice, error := ioutil.ReadFile(filename)

	if error != nil {
		fmt.Println("Something went wrong: ", error)
		os.Exit(1)
	}

	cardsAsString := string(byteSlice)
	cardsSliceOfStrings := strings.Split(cardsAsString, ",")

	cardDeck := deck(cardsSliceOfStrings)

	return cardDeck
}
