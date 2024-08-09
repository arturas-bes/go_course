package main

import (
	"os"
	"testing"
)

//Here we are creating a test function for the newDeck function
func TestNewDeck(t *testing.T) {
	d := newDeck()

	//Here we are checking if the length of the deck is 16
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	//Here we are checking if the first card is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])

	}

	//Here we are checking if the last card is Four of Clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but got %v", d[len(d)-1])

	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//Here we are removing the file _decktesting
	os.Remove("_decktesting")
	d := newDeck()
	//Here we are saving the deck of cards to a file
	d.saveToFile("_decktesting")

	//Here we are creating a new deck of cards from the file _decktesting
	ld := newDeckFromFile("_decktesting")

	//Here we are checking if the length of the deck is 16
	if len(ld) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(ld))
	}

	//Here we are removing the file _decktesting
	os.Remove("_decktesting")
}
