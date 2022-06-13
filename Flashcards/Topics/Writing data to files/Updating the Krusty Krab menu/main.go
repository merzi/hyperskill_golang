package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// DO NOT DELETE! This opens the 'galley_grub.txt' file in read-write|append mode
	file, err := os.OpenFile("galley_grub.txt", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	// write "Kelp Shake $2.00" to the 'file' variable below with the fmt.Fprintln() function
	_, err = fmt.Fprintln(file, "Kelp Shake $2.00")
	if err != nil {
		log.Println(err)
	}
}
