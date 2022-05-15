package main

import "fmt"

// What additional syntax does a function parameter require
// to be able to take multiple arguments?
func divide(nums ...float64) float64 {
	total := 1.0
	for _, num := range nums {
		total /= num
	}
	return total
}

func main() {
	// DO NOT change any code within the 'main' function!
	fmt.Println(divide())
	fmt.Println(divide(1, 2))
	fmt.Println(divide(0, 1))
}
