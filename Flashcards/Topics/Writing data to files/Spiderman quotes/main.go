package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	quote := readQuote() // today's quote
	file, err := os.Create("quote.txt")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: write the quote to a file
	file.WriteString(quote)
}

func readQuote() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	_ = scanner.Scan()
	return scanner.Text()
}
