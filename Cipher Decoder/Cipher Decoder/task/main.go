package main

import (
	"fmt"
	"math"
	"math/rand"
)

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

	var g, p, A int
	var message string
	fmt.Scanf("g is %d and p is %d", &g, &p)
	fmt.Println("OK")
	b := rand.Intn(p)
	a := rand.
	B := calculateCipher(g, p, b)
	fmt.Scanf("A is %d", &A)
	bSecret := computeSharedSecret(B, A, p)
	aSecret := computeSharedSecret(A, B, p)
	fmt.Printf("B is %d\n", bSecret)
	fmt.Println(encryptCaesar("Will you marry me?", aSecret))
	fmt.Scanln(&message)
	if decryptCaesar(message, bSecret) == "Yeah, okay!" {
		message = "Great!"
	} else {
		message = "What a pity!"
	}
	fmt.Println(encryptCaesar(message, aSecret))
}

func calculateCipher(g int, p int, secret int) int {
	return int(math.Pow(float64(g), float64(secret))) % p
}

func computeSharedSecret(cipher int, foreignSecret int, p int) int {
	return int(math.Pow(float64(cipher), float64(foreignSecret))) % p
}

func encryptCaesar(plain string, shift int) string {
	var encrypted string

	for _, value := range []byte(plain) {
		if isIgnoredChar(string(value)) {
			encrypted += string(value)
			continue
		}
		encrypted += shiftChar(value, shift)
	}
	return encrypted
}

func decryptCaesar(encrypted string, shift int) string {
	var decrypted string
	for _, value := range []byte(encrypted) {
		if isIgnoredChar(string(value)) {
			decrypted += string(value)
			continue
		}
		decrypted += shiftChar(value, shift*-1)
	}

	return decrypted
}

func shiftChar(char byte, shift int) string {
	shift = (shift%26 + 26) % 26 // [0, 25]
	var a int
	switch {
	case 'a' <= rune(char) && rune(char) <= 'z':
		a = 'a'
	case 'A' <= rune(char) && rune(char) <= 'Z':
		a = 'A'
	default:
		return string(char)
	}
	return string(byte(a + ((int(char)-a)+shift)%26))
}

func isIgnoredChar(char string) bool {
	ignoredChars := []string{
		"!",
		"?",
		":",
		",",
		".",
		";",
		" ",
		"-",
		"_",
	}
	return sliceContains(ignoredChars, char)
}

func sliceContains(slice []string, str string) bool {
	for _, value := range slice {
		if value == str {
			return true
		}
	}
	return false
}
