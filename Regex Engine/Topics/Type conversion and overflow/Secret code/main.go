package main

import "fmt"

func main() {
	var s string
	var shift int
	fmt.Scan(&s, &shift)

	for _, ch := range s {
		fmt.Print(string(int(ch) + shift)) // add the shift value to the character value
	}
}
