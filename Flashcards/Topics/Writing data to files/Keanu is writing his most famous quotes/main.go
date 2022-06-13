package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// DO NOT DELETE! This opens the 'keanu_quotes.txt' file in read-write mode
	file, err := os.OpenFile("keanu_quotes.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	quotes := []string{"You're breathtaking!",
		"Wake up Samurai! We have a city to burn.",
		"Lose? I don't lose; I win! I'm a lawyer, that's my job, that's what I do!",
	}

	for _, line := range quotes { // iterate over each line of the 'quotes' slice here
		_, err := fmt.Fprintln(file, line) // // write each line of the 'quotes' slice of strings to 'file' here
		if err != nil {
			log.Fatal(err) // exit if we have an unexpected error
		}
	}
}
