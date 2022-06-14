package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// DO NOT delete! - This code block takes as an input the values for the 'movies' map:
	var title string
	var releaseYear int
	var rating float64

	fmt.Scanln(&title, &releaseYear, &rating)

	songs := map[string]interface{}{
		"title":       title,
		"releaseYear": releaseYear,
		"rating":      rating,
	}

	// Call the json.MarshalIndent() function below and pass
	// The correct 'prefix' and 'indent' following Bart's request:
	moviesJson, err := json.MarshalIndent(songs, "\t", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(moviesJson))
}
