package main

import (
    "fmt"
    "io/ioutil"
    "math/rand"
    "os"
    "strings"
    "time"
)

//Here we are creating a new type of deck which is a slice of strings
type deck []string

func newDeck() deck {
    //Here we are creating a new deck of cards
    cards := deck{}

    //Here we are creating two slices of strings
    cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
    cardValues := []string{"Ace", "Two", "Three", "Four"}

    //Here we are using a nested loop to create a deck of cards
    for _, suit := range cardSuits {
        for _, value := range cardValues {
            cards = append(cards, value+" of "+suit)
        }
    }
    return cards
}

//Here we are creating a receiver function for the deck type
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}

//Here we are creating a function that returns two decks
func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}

//Here we are creating a receiver function for the deck type
func (d deck) toString() string {
    return strings.Join([]string(d), ",")
}

//Here we are creating a receiver function for the deck type
func (d deck) saveToFile(filename string) error {
    return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//Here we are creating a function that reads a file and returns a deck
func newDeckFromFile(filename string) deck {
    //Here we are reading the file and returning a byte slice and an error
    bs, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    //Here we are converting the byte slice to a string and then splitting it into a slice of strings
    s := strings.Split(string(bs), ",")
    return deck(s)
}

//Here we are creating a receiver function for the deck type
func (d deck) shuffle() {
    //Here we are creating a new source for the random number generator
    source := rand.NewSource(time.Now().UnixNano())
    //Here we are creating a new random number generator
    r := rand.New(source)

    //Here we are using the Fisher-Yates shuffle algorithm to shuffle the deck
    for i := range d {
        //Here we are generating a random number between 0 and the length of the deck
        newPosition := r.Intn(len(d) - 1)

        //Here we are swapping the current card with the card at the random position
        d[i], d[newPosition] = d[newPosition], d[i]
    }
}