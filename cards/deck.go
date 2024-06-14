package main

import "fmt"

//We gonna create deck type here which is a slice of strings

// In case we want to compile both files we should use go run "use both files needed as main is executable and deck contains object" 
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//There was vriables i, j, but we don't need that here so we replaced them with _ so compiler does not complain 
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//Receiver function "cards" refers to actual data in receiver type in go self, this  are not used
//Usually names of actual data are used by one or two letters of that type in this case is d as deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}