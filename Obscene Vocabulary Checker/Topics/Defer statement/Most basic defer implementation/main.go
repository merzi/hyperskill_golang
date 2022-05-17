package main

import "fmt"

func main() {
	var firstName string
	var lastName string

	// delete or add new defer statements below!
	fmt.Scanln(&firstName)
	fmt.Scanln(&lastName)

	defer fmt.Printf("%s %s", firstName, lastName)
	fmt.Printf("%s,\n", lastName)
}
