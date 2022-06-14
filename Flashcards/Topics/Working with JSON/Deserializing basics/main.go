package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// DO NOT delete! - This code block takes as an input the values for the
	// 'carsJson' and 'dogsJson' variables:
	scanner := bufio.NewScanner(os.Stdin)

	var carsJson string
	scanner.Scan()
	carsJson = scanner.Text()

	var dogsJson string
	scanner.Scan()
	dogsJson = scanner.Text()

	var cars map[string]interface{}
	err := json.Unmarshal([]byte(carsJson), &cars)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cars)

	var dogs []string
	err = json.Unmarshal([]byte(dogsJson), &dogs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dogs)
}
