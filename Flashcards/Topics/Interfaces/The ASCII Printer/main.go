package main

import (
	"fmt"
	"log"
	"os"
)

type File string

// Do not change the contents of the PrintAscii() method!
func (f File) PrintAscii() {
	b, err := os.ReadFile(string(f))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

// Create the AsciiPrinter interface with the PrintAscii() method below:
type AsciiPrinter interface {
	PrintAscii()
}

func main() {
	// Create the variable 'a' of the AsciiPrinter interface type below:
	var a AsciiPrinter

	// Open and read the file "ascii_art.txt" with the 'a' AsciiPrinter interface:
	a = File("ascii_art.txt")

	// Call the PrintAscii() method on the 'a' AsciiPrinter interface below:
	a.PrintAscii()
}
