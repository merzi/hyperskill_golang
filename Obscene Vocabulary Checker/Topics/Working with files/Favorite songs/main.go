package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Write the code below to open "songs.txt" and store it in the 'file' variable:
	file, err := os.Open("songs.txt")

	// Implement error handling below:
	if err != nil {
		log.Fatal(err) // exit if we have an error
	}

	// DO NOT delete this code block - it outputs the contents of the file after opening it
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	file.Close() // always remember to close the file!
}
