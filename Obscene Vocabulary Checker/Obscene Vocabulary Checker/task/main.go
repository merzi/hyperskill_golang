package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// write your code here
	var forbiddenTextFile, sentence string
	var forbiddenWords []string
	_, err := fmt.Scanln(&forbiddenTextFile)
	if err != nil {
		return
	}

	forbiddenWords = getForbiddenWords(forbiddenTextFile)
	if forbiddenWords == nil {
		return
	}

	running := true
	for running {
		_, err = fmt.Scanln(&sentence)
		if err != nil {
			fmt.Println("invalid input!")
			continue
		}

		if sentence == "exit" {
			running = false
			fmt.Println("Bye!")
			continue
		}

		fmt.Println(censorSentence(sentence, forbiddenWords))
	}
}

func censorSentence(sentence string, forbiddenWords []string) string {
	var outputSlice []string
	var word string
	scanner := bufio.NewScanner(strings.NewReader(sentence))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		if sliceContains(forbiddenWords, scanner.Text(), false) > -1 {
			word = strings.Repeat("*", utf8.RuneCountInString(scanner.Text()))
		} else {
			word = scanner.Text()
		}
		outputSlice = append(outputSlice, word)
	}

	return strings.Join(outputSlice, " ")
}

func getForbiddenWords(fileName string) []string {
	var forbiddenWords []string
	fileOpener, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer func(fileOpener *os.File) {
		err := fileOpener.Close()
		if err != nil {

		}
	}(fileOpener)

	scanner := bufio.NewScanner(fileOpener)

	for scanner.Scan() {
		forbiddenWords = append(forbiddenWords, scanner.Text())
	}

	return forbiddenWords
}

func sliceContains(slice []string, word string, caseSensitive bool) int {
	for key, forbiddenWord := range slice {
		if caseSensitive && forbiddenWord == word {
			return key
		} else if caseSensitive == false && strings.ToLower(forbiddenWord) == strings.ToLower(word) {
			return key
		}
	}

	return -1
}
