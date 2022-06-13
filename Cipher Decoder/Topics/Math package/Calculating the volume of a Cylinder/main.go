package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Scanln(&radius)

	var height float64
	fmt.Scanln(&height)

	// Using the functions from the math package create the formula
	// to calculate the volume of a cylinder below:
	volume := math.Pi * math.Pow(radius, 2) * height

	fmt.Printf("%.2f", volume)
}
