package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("galley_grub.txt") // open the file here with the os.ReadFile() function
	if err != nil {
		log.Fatal(err) // exit if we have an unexpected error
	}
	// print the contents of the file here as a string
	fmt.Print(string(file))
}
