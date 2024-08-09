package main

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//Here we are iterating over the numbers slice
	for _, nunumbers := range numbers {
		//Here we are checking if the number is even or odd
		if nunumbers % 2 == 0 {
			println(nunumbers, "is even")
		} else {
			println(nunumbers, "is odd")
		}
	}
}