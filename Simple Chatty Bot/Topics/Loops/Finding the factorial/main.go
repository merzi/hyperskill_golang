package main

import "fmt"

func main() {
	// put your code here
	var integer, factVal int
	fmt.Scan(&integer)
	factVal = 1
	if integer > 0 {
		for i := 1; i <= integer; i++ {
			factVal *= i // mismatched types int64 and int
		}
	}

	fmt.Println(factVal)
}
