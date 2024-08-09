package main

import "fmt"

//here we are defining a new type of bot which is an interface
//Iterfaces cannot be used as generic types(Go lang does not like it)
//Iterfaces are implicit, we do not need to explicitly say that we are implementing an interface
//Interfaces are a contract to help us manage types
type bot interface {
	getGreetng() string //Here we can have multiple methods in the interface, multiple arguments and return types
}

//Here we are defining a new type of englishBot which is a struct
type enligshBot struct{}
type spanishBot struct{}

func main() {
	eb := enligshBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// Here interface is used to define a common method for different types of bots
func printGreeting(b bot) {
	fmt.Println(b.getGreetng())
}

// Here eb can be removed it is optional if not actualy used
func (eb enligshBot) getGreetng() string {
	//Very custom logic for generating an english greeting
	return "Hi there!"
}

func (sb spanishBot) getGreetng() string {
	//Very custom logic for generating a spanish greeting
	return "Hola!"
}
