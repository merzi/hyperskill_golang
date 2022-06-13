package main

import "fmt"

func main() {
	var a, b int8
	fmt.Scan(&a, &b)

	if int(a+b) != int(a)+int(b) {
		fmt.Println("+")
	}

	// add the same checks for "-" and "*"
	if int(a-b) != int(a)-int(b) {
		fmt.Println("-")
	}

	if int(a*b) != int(a)*int(b) {
		fmt.Println("*")
	}
}
