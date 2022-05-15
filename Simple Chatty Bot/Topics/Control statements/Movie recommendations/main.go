package main

import "fmt"

func main() {
	var age int
	fmt.Scanf("%d", &age)

	// Code your switch or if...else-if statement here.
	switch true {
	case age <= 14:
		fmt.Println("Toy Story 4")
		break
	case age <= 18:
		fmt.Println("The Matrix")
		break
	case age <= 25:
		fmt.Println("John Wick")
		break
	case age <= 35:
		fmt.Println("Constantine")
		break
	default:
		fmt.Println("Speed")
	}
}
