package main

import (
	"fmt"
	"sort"
)

func main() {
	champions := []string{"Jax", "Mordekaiser", "Singed", "Rumble", "Irelia"}

	// Write the required logic below to sort the 'champions' slice based on
	// the shortest name to longest name!
	sort.Slice(champions, func(i, j int) bool {
		return len(champions[i]) < len(champions[j])
	})

	fmt.Println(champions)
}
