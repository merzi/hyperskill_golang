package main

import "fmt"

func main() {
	// write your code here
	/*	p := 23
		base := 5

		aliceSecret := 4
		bobSecret := 3

		A := int(math.Pow(float64(base), float64(aliceSecret))) % p
		B := int(math.Pow(float64(base), float64(bobSecret))) % p

		sAlice := int(math.Pow(float64(B), float64(aliceSecret))) % p
		sBob := int(math.Pow(float64(A), float64(bobSecret))) % p
	*/

	var g, p int
	fmt.Scanf("g is %d and p is %d", &g, &p)
	fmt.Printf("g=%d and p=%d", g, p)

}
