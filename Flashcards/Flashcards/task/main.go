package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	// write your code here
	cards := map[string]string{}

	var term, definition, answer string

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	term = strings.TrimSpace(line)
	line, _ = reader.ReadString('\n')
	definition = strings.TrimSpace(line)

	cards = addCard(term, definition, cards)
	line, _ = reader.ReadString('\n')
	answer = strings.TrimSpace(line)

	solveCard(term, answer, cards)
}

func addCard(term string, definition string, cards map[string]string) map[string]string {
	cards[term] = definition

	return cards
}

func solveCards(cards map[string]string) {
	var term, definition string
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	term = strings.TrimSpace(line)
	line, _ = reader.ReadString('\n')
	definition = strings.TrimSpace(line)

	//	fmt.Println(cards[term])
	solveCard(term, definition, cards)
}

func solveCard(term string, definition string, cards map[string]string) bool {
	if cards[term] == definition {
		fmt.Println("Your answer is right!")
		return true
	}
	fmt.Println("Your answer is wrong...")
	return false
}

func getCards() map[string]string {
	return map[string]string{
		"purchase":         "buy",
		"buy":              "purchase",
		"cos'(x)":          "-sin(x)",
		"-sin(x)":          "cos'(x)",
		"a purring animal": "lion",
		"lion":             "a purring animal",
	}
}

func playRandomCards() {
	cardFront, cardBack := pickRandomCard(getCards())
	fmt.Println("Card:")
	fmt.Println(cardFront)
	fmt.Println("Definition:")
	fmt.Println(cardBack)
}

func pickRandomCard(cards map[string]string) (string, string) {
	k := rand.Intn(len(cards) - 1)
	i := 0
	for key, value := range cards {
		if i == k {
			return key, value
		}
		i++
	}

	return "unknown", "unknown"
}
