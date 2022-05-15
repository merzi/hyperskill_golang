package main

import (
	"fmt"
)

func main() {
	// put your code here
	var digits, reversed string
	fmt.Scanln(&digits)

	for i := len(digits) - 1; 0 <= i; i-- {
		reversed += string(digits[i])
	}

	fmt.Println(reversed)
}
