package main

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, nunumbers := range numbers {
		if nunumbers % 2 == 0 {
			println(nunumbers, "is even")
		} else {
			println(nunumbers, "is odd")
		}
	}
}