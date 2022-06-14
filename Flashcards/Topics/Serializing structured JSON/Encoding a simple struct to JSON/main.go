package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// DO NOT delete or modify the FootballPlayer struct implementation:
type FootballPlayer struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      int    `json:"age"`
	Team     string `json:"team"`
	Position string `json:"position"`
}

func main() {
	// DO NOT delete! - This code block takes as an input the values for the 'player' struct:
	var name, lastName, team, position string
	var age int
	fmt.Scanln(&name, &lastName, &age, &team, &position)

	player := FootballPlayer{
		Name:     name,
		LastName: lastName,
		Age:      age,
		Team:     team,
		Position: position,
	}

	// Write the code below to properly encode the 'player' struct into the 'playerJson' variable:
	playerJson, err := json.Marshal(player)
	if err != nil {
		log.Fatal(err)
	}

	// Print the serialized 'playerJson' below; remember to cast it as a 'string'!
	fmt.Println(string(playerJson))
}
