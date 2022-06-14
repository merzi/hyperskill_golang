package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// DO NOT delete! - This code block takes as an input the values for the 'songs' map:
	var songName string
	var songDuration string
	var songReleaseYear int

	fmt.Scanln(&songName, &songDuration, &songReleaseYear)

	// What type in Go allows us to refer to "any" type!?
	// Correct the type of the values of the songs map below to accept "any" type of values!
	songs := map[string]any{
		"name":        songName,
		"duration":    songDuration,
		"releaseYear": songReleaseYear,
	}

	// Do NOT delete! - This code block encodes 'songs' to 'songsJson'
	// And then prints the encoded result as a string!
	songsJson, err := json.Marshal(songs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(songsJson))
}
