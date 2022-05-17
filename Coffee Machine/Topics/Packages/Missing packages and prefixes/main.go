package main

import (
	// write the missing imports here
	"fmt"
	"math"
)

func main() {
	// just add the missing prefixes to the functions below.
	var angle float64
	fmt.Scanf("%f", &angle)

	angle = angle * (math.Pi / 180)

	fmt.Println(math.Sin(angle) - math.Cos(angle))
}
