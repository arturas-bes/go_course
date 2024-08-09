package main

import "fmt"

func main() {
	//Here we are creating a slice of strings
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	//Here we are passing the slice to the updateSlice function
	updateSlice(mySlice)

	//Here we are printing the slice
	fmt.Println(mySlice)

	//Idea is that slice points to the same underlying array so if we update the slice it will update the underlying array on the same memory location
 }

 func updateSlice(s []string) {
	//Here we are updating the first element of the slice
	s[0] = "Bye"
 }