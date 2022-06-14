package main

import (
	"fmt"
	"sort"
)

type Team struct {
	Name   string
	Points int
}

func main() {
	teams := []Team{
		{"Borussia Dortmund", 64},
		{"Bayern Münich", 78},
		{"RB Leipzig", 65},
		{"Vfl Wolfsburg", 61},
	}

	// Implement the logic within the sort.Slice() function below to sort the teams slice!
	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Points > teams[j].Points
	})

	// Do not delete the output line!
	fmt.Println(teams)
}
