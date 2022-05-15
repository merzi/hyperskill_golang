package main

import "fmt"

func formatUsernameText(username string) string { // function name, parameters and return type here {
	// write the function code block here

	return fmt.Sprintf("%s is learning how to call functions!", username)
}

func main() {
	// Do not change this two lines of code
	var userName string
	fmt.Scanf("%s", &userName)

	// call the function directly, or within a fmt.Println statement here.
	fmt.Println(formatUsernameText(userName))
}
