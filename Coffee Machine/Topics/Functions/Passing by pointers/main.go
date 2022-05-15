package main

import "fmt"

// Change the 'updateRapper' function type of parameters and the type of variables within it:
func updateRapper(fName *string, lName *string) {
	*fName = "Kanye" // do not change the name!
	*lName = "West"  // do not change the last name!
}

func main() {
	// DO NOT delete! -- This code block takes as an input the name/last name of the rapper
	var firstName, lastName string
	fmt.Scanln(&firstName, &lastName)

	fmt.Println("My favorite rapper was:", firstName, lastName)

	updateRapper(&firstName, &lastName)

	fmt.Println("My new favorite rapper is:", firstName, lastName)
}
