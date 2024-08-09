package main

//Here we are importing the fmt package
func main() {
	//Here we are creating a new deck of cards
	cards := newDeck()
	//Here we are shuffling the deck of cards
	cards.shuffle()
	//Here we are printing the deck of cards
	cards.print()
}