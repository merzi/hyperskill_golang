package main

import (
	"fmt"
	"strings"
)

func main() {
	// 	write your code here
	var regexString, inputString string
	fmt.Scan(&regexString)

	inputString = strings.Split(regexString, "|")[1]
	regexString = strings.Split(regexString, "|")[0]

	switch true {
	case inputString == regexString:
		fmt.Println(true)
		break
	case regexString == ".":
		fmt.Println(true)
		break
	case regexString == "":
		fmt.Println(true)
		break
	case regexString != inputString:
		fmt.Println(false)
		break
	case regexString == "" && inputString == "":
		fmt.Println(true)
		break
	case regexString != "" && inputString == "":
		fmt.Println(false)
		break
	}

}
