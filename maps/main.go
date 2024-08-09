package main

func main() {

	// //Here we are declaring a map of colors with a key of string and a value of string
	// var colors map[string]string

	//Here we are creating a map of colors using the map function
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//Here we are creating a map of colors using the make function
	// colors := make(map[string]string)

	//Here we are adding key value pairs to the map
	// colors["white"] = "#ffffff"

	//Here we are deleting a key value pair from the map
	// delete(colors, "white")

	
	printMap(colors)

	//The main differences between maps and structs are that maps are used to store a collection of related properties and structs are used to store a collection of properties that describe a single thing
	//The most useful example for struct and maps would be to use a struct to describe a person and a map to describe a collection of people
	}

func printMap(c map[string]string) {
	//Here we are iterating over the map colors, the key is color and the value is hex
	for color, hex := range c {
		println("Hex code for", color, "is", hex)
	}
}