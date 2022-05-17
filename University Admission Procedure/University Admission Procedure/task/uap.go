// write your code here
package main

import "fmt"

func main() {
	var number1, number2, number3 int

	fmt.Scanf("%d", &number1)
	fmt.Scanf("%d", &number2)
	fmt.Scanf("%d", &number3)

	meanScore := (float64(number1) + float64(number2) + float64(number3)) / 3
	fmt.Println(fmt.Sprintf("%.2f", meanScore))
	if meanScore >= 60.0 {
		fmt.Println("Congratulations, you are accepted!")
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
	}

}
