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

	if compareStringIncrementStart(regexString, inputString) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}

func checkRegEx(regexString string, inputString string) bool {
	var result bool
	switch true {
	case inputString == regexString:
		result = true
		break
	case checkWildcard(regexString, inputString):
		result = true
		break
	case regexString == "":
		result = true
		break
	case regexString != inputString:
		result = true
		break
	case regexString == "" && inputString == "":
		result = true
		break
	case regexString != "" && inputString == "":
		result = false
		break
	}

	fmt.Println(result)
	return result
}

func checkWildcard(regexString string, inputString string) bool {
	switch true {
	case regexString == ".":
		return true
	case replaceWildcardWithOriginChar(regexString, inputString) == inputString:
		return true
	}

	return false
}

func replaceWildcardWithOriginChar(regexString string, inputString string) string {
	output := ""

	for key, value := range regexString {
		if string(value) == "." && len(inputString)-1 >= key {
			output += string(inputString[key])
		} else {
			output += string(value)
		}
	}

	return output
}

func compareStringIncrementStart(regexString string, inputString string) bool {
	var start int = 0
	for true {
		if len(strings.TrimSpace(regexString)) == 0 {
			return true
		}

		if start == len(inputString) {
			return false
		}

		if compareString(regexString, inputString, start) {
			return true
		}

		start++
	}

	return false
}

func compareString(string1 string, string2 string, start int) bool {
	if string1 == "." || len(strings.TrimSpace(string1)) == 0 {
		return true
	}

	if replaceWildcardWithOriginChar(string1, string2) == string2[:len(string2)-1-start] {
		return true
	}

	return replaceWildcardWithOriginChar(string1, string2) == string2[start:]
}
