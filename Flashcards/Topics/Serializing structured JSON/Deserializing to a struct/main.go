package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// DO NOT delete or modify the FootballTeams struct implementation:
type FootballTeams struct {
	League string `json:"league"`
	Clubs  []struct {
		Name    string `json:"name"`
		Code    string `json:"code"`
		Stadium string `json:"stadium"`
	} `json:"clubs"`
}

func main() {
	// DO NOT delete! - This block takes as an input the JSON object for the 'teamsJson' variable:
	scanner := bufio.NewScanner(os.Stdin)

	var teamsJson string
	scanner.Scan()
	teamsJson = scanner.Text()

	// Write the required code below to properly decode the 'teamsJson' JSON object to the 'teams' struct:
	var teams FootballTeams
	err := json.Unmarshal([]byte(teamsJson), &teams)
	if err != nil {
		log.Fatal(err)
	}

	// Print the decoded 'teams' struct below -- do not change the "%+v" verb!
	fmt.Printf("%+v", teams)
}
