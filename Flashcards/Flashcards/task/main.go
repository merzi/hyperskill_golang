package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// write your code here
	/*	cards := map[string]string{}

		var term, definition, answer string

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		term = strings.TrimSpace(line)
		line, _ = reader.ReadString('\n')
		definition = strings.TrimSpace(line)

		cards = addCard(term, definition, cards)
		line, _ = reader.ReadString('\n')
		answer = strings.TrimSpace(line)

		solveCard(term, answer, cards) */
	var c Cards
	var numCards int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input the number of cards: ")
	scanner.Scan()
	numCards, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	c.addAmountCards(numCards)
	c.solveAll()
}

type Card struct {
	term       string
	definition string
}

type Cards struct {
	Card []Card
}

func (c *Cards) addAmountCards(amount int) bool {
	oldAmount := len(c.Card)
	var term, definition string
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < amount; i++ {
		fmt.Printf("The term for card #%d:\n", i+1)
		scanner.Scan()
		term = strings.TrimSpace(scanner.Text())

		if c.checkTermAlreadyExists(term) {
			fmt.Printf("The term \"%s\" already exists. Try again:", term)
			amount++
			continue
		}

		fmt.Printf("The definition for card #%d:\n", i+1)
		scanner.Scan()
		definition = strings.TrimSpace(scanner.Text())
		if c.checkDefinitionAlreadyExists(definition) {
			fmt.Printf("The definition \"%s\" already exists. Try again:", definition)
			amount++
			continue
		}

		c.addCard(term, definition)
		term = ""
		definition = ""
	}

	return len(c.Card) == oldAmount+amount
}

func (c *Cards) checkTermAlreadyExists(term string) bool {
	for _, card := range c.Card {
		if card.term == term {
			return true
		}
	}

	return false
}

func (c *Cards) checkDefinitionAlreadyExists(definition string) bool {
	for _, card := range c.Card {
		if card.definition == definition {
			return true
		}
	}

	return false
}

func (c *Cards) findTermByDefinition(definition string) (card Card, err error) {
	for _, card := range c.Card {
		if card.definition == definition {
			return card, err
		}
	}

	err = errors.New("definition not found!")
	return Card{}, err
}

func (c *Cards) addCard(term string, definition string) {
	c.Card = append(c.Card, Card{term: term, definition: definition})
}

func (c *Cards) deleteCardByTermDefinition(term string, definition string) bool {
	for key, currentCard := range c.Card {
		if currentCard.term == term && currentCard.definition == definition {
			return c.deleteCard(key)
		}
	}

	return false
}

func (c *Cards) deleteCard(cardNumber int) bool {
	oldLength := len(c.Card)
	c.Card = append(c.Card[:cardNumber], c.Card[cardNumber+1:]...)
	return len(c.Card)+1 == oldLength
}

func (c *Card) updateCard(term string, definition string) bool {
	c.term = term
	c.definition = definition

	return c.term == term && c.definition == definition
}

func (c *Cards) updateCardByOldValues(term string, definion string, oldTerm string, oldDefinition string) bool {
	for _, currendCard := range c.Card {
		if currendCard.term == oldTerm && currendCard.definition == oldDefinition {
			return currendCard.updateCard(term, definion)
		}
	}

	return false
}

func (c *Cards) updateCard(term string, definition string, cardNumber int) bool {
	c.Card[cardNumber].updateCard(term, definition)

	return c.Card[cardNumber].term == term && c.Card[cardNumber].definition == definition
}

func (c *Cards) solveAll() bool {
	var advisedSolution string
	var result bool
	result = true
	for _, currentCard := range c.Card {
		fmt.Printf("Print the definition of \"%s\"\n", currentCard.term)
		_, err := fmt.Scan(&advisedSolution)
		if err != nil {
			return false
		}
		if currentCard.solve(advisedSolution) {
			fmt.Println("Correct!")
		} else {
			otherCard, err := c.findTermByDefinition(advisedSolution)
			if err == nil {
				fmt.Printf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".\n",
					currentCard.getDefinition(),
					otherCard.term)
			} else {
				fmt.Printf("Wrong. The right answer is \"%s\".\n", currentCard.getDefinition())
			}
			result = false
		}
	}

	return result
}

func (c *Cards) solve(term string, definition string) bool {
	for _, currendCard := range c.Card {
		if currendCard.term == term {
			return currendCard.solve(definition)
		}
	}

	return false
}

func (c *Card) solve(definition string) bool {
	return c.definition == definition
}

func (c *Card) getDefinition() string {
	return c.definition
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
