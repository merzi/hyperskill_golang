package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	exportTo   *string
	importFrom *string
)

func init() {
	exportTo = flag.String("export_to", "", "filename for the auto safe file")
	importFrom = flag.String("import_from", "", "filename for the auto load file")
}

func main() {
	// write your code here
	flag.Parse()
	*exportTo = strings.TrimSpace(*exportTo)
	*importFrom = strings.TrimSpace(*importFrom)

	var p = program{Cards: Cards{Scanner: bufio.NewScanner(os.Stdin)},
		AutoLoadFile: *importFrom, AutoSafeFile: *exportTo}

	p.Cards.print("Input the program version (first, second, third, fourth, final)[first is default]: ")
	switch p.Cards.getInput() {
	case "first":
		p.firstStage()
		break
	case "second":
		p.secondStage()
		break
	case "third":
		p.thirdStage()
		break
	case "fourth":
		p.fourthStage()
		break
	case "final":
		p.finalStage()
		break
	default:
		p.firstStage()
	}
}

type program struct {
	Cards        Cards
	AutoLoadFile string
	AutoSafeFile string
}

func (p *program) firstStage() {
	playRandomCards()
}

func (p *program) secondStage() {
	solveCards(getCards())
}

func (p *program) thirdStage() {
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

func (p *program) fourthStage() {
	var numCards int
	p.Cards.print("Input the number of cards: ")
	numCards, err := strconv.Atoi(p.Cards.getInput())
	if err != nil {
		return
	}

	p.Cards.addAmountCards(numCards)
	p.Cards.solveAll()
}

func (p *program) finalStage() {
	if p.checkFileExists(p.AutoLoadFile) {
		p.Cards.importCards(p.AutoLoadFile)
		if len(p.Cards.CardStack) > 0 {
			p.Cards.print(fmt.Sprintf("%d cards have been loaded.", len(p.Cards.CardStack)))
		} else {
			p.Cards.print("auto import failed!")
		}
	}

	for true {
		if !p.menu() {
			break
		}
	}
}

func (p *program) menu() bool {
	p.Cards.print("Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats): ")
	switch p.Cards.getInput() {
	case "add":
		term := p.Cards.requestCard()
		definition := p.Cards.requestCardDefinition()
		if p.Cards.addCard(term, definition) {
			p.Cards.print(fmt.Sprintf("The pair (\"%s\":\"%s\") has been added.", term, definition))
		} else {
			p.Cards.print(fmt.Sprintf("The pair (\"%s\":\"%s\") has not been added.", term, definition))
		}
		break
	case "remove":
		p.Cards.print("Which Cards?")
		input := p.Cards.getInput()
		if p.Cards.deleteCardByTerm(input) {
			p.Cards.print("The card has been removed.")
		} else {
			p.Cards.print(fmt.Sprintf("Can't remove \"%s\": there is no such card.", input))
		}
		break
	case "import":
		p.Cards.print("File name: ")
		input := p.Cards.getInput()
		if !p.checkFileExists(input) {
			break
		}

		amount := p.Cards.importCards(input)
		if amount >= 0 {
			p.Cards.print(fmt.Sprintf("%d cards have been loaded.", amount))
		} else {
			p.Cards.print("import failed!")
		}
		break
	case "export":
		p.Cards.print("File name: ")
		input := p.Cards.getInput()
		amount := p.Cards.exportCards(input)
		if amount >= 0 {
			p.Cards.print(fmt.Sprintf("%d cards have been saved.", amount))
		} else {
			p.Cards.print("export failed!")
		}
		break
	case "ask":
		if len(p.Cards.CardStack) < 1 {
			p.Cards.print("No Cards found! Import or add Cards first!")
			break
		}

		p.Cards.print("How many times to ask? ")
		input := p.Cards.getInput()
		numCards, err := strconv.Atoi(input)
		if err != nil {
			p.Cards.print(fmt.Sprintf("can't convert \"%s\" to integer!", input))
			break
		}
		p.Cards.ask(numCards)
		break
	case "exit":
		if p.checkFileName(p.AutoSafeFile) {
			if p.Cards.exportCards(p.AutoSafeFile) > 0 {
				p.Cards.print(fmt.Sprintf("%d cards have been saved.", len(p.Cards.CardStack)))
			} else {
				p.Cards.print("export failed")
			}
		}

		p.Cards.print("Bye bye!")
		return false
	case "log":
		p.Cards.print("File name: ")
		err := p.Cards.saveLogToFile(p.Cards.getInput())
		if err != nil {
			p.Cards.print("Log export failed!")
			break
		}

		p.Cards.print("The Log has been saved.")
		break
	case "hardest card":
		count, cards := p.Cards.getHardestCard()
		if count == 0 || len(cards) < 1 {
			p.Cards.print("There are no cards with errors.")
		} else if len(cards) > 1 {
			p.Cards.print(fmt.Sprintf("The hardest cards are \"%d\". You have %d errors answering them.",
				strings.Join(cards, "\", \""), count))
		} else {
			p.Cards.print(fmt.Sprintf("The hardest card is \"%s\". You have %d errors answering it.",
				cards[0], count))
		}

		break

	case "reset stats":
		if p.Cards.reset() {
			p.Cards.print("Card statistics have been reset.")
		} else {
			p.Cards.print("Card statistics haven't been reset.")
		}
		break
	default:
		p.Cards.print("unknown command!")
		break
	}
	p.Cards.print("")
	return true
}

func (p *program) checkFileExists(fileName string) bool {
	if !p.checkFileName(fileName) {
		return false
	}

	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		p.Cards.print("File not found.")
		return false
	} else if err != nil {
		p.Cards.print("import file failure!")
		return false
	}

	return true
}

func (p *program) checkFileName(fileName string) bool {
	fileNameParts := strings.Split(fileName, ".")

	return len(fileNameParts) > 1 &&
		len(strings.TrimSpace(fileNameParts[0])) > 0 &&
		len(strings.TrimSpace(fileNameParts[len(fileNameParts)-1])) > 0
}

type Card struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
	Failures   int    `json:"failed_tries"`
	Tries      int    `json:"tries"`
}

type Cards struct {
	CardStack []Card
	Log       []string
	Scanner   *bufio.Scanner
}

func (c *Cards) getInput() string {
	c.Scanner.Scan()
	c.addLogEntry(c.Scanner.Text())

	return c.Scanner.Text()
}

func (c *Cards) print(output string) bool {
	fmt.Println(output)

	return c.addLogEntry(output)
}

func (c *Cards) addLogEntry(entry string) bool {
	c.Log = append(c.Log, entry)

	return c.Log[len(c.Log)-1] == entry
}

func (c *Cards) saveLogToFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	write := bufio.NewWriter(file)
	for _, value := range c.Log {
		_, err := fmt.Fprintln(write, value)
		if err != nil {
			return err
		}
	}

	return write.Flush()
}

func (c *Cards) exportCards(fileName string) int {
	return c.exportToFile(fileName)
}

func (c *Cards) exportToFile(fileName string) int {
	jsonObject, err := json.Marshal(c.CardStack)
	if err != nil {
		return -1
	}

	err = os.WriteFile(fileName, jsonObject, 0644)
	if err != nil {
		return -1
	}
	return len(c.CardStack)
}

func (c *Cards) importCards(fileName string) int {
	return c.importFromFile(fileName)
}

func (c *Cards) importFromFile(fileName string) int {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return -1
	}

	err = json.Unmarshal(file, &c.CardStack)
	if err != nil {
		return -1
	}

	return len(c.CardStack)
}

func (c *Cards) addAmountCards(amount int) bool {
	oldAmount := len(c.CardStack)

	for i := 0; i < amount; i++ {
		c.addCard(c.requestTerm(i), c.requestDefinition(i))
	}

	return len(c.CardStack) == oldAmount+amount
}

func (c *Cards) requestTerm(num int) string {
	var term string
	c.print(fmt.Sprintf("The term for card #%d: ", num+1))
	for true {
		term = strings.TrimSpace(c.getInput())

		if !c.checkTermAlreadyExists(term) {
			return term
		}
		c.print(fmt.Sprintf("The term \"%s\" already exists. Try again: ", term))
	}

	return term
}

func (c *Cards) requestCard() string {
	var term string
	c.print("The card: ")
	for true {
		term = strings.TrimSpace(c.getInput())

		if !c.checkTermAlreadyExists(term) {
			return term
		}
		c.print(fmt.Sprintf("The card \"%s\" already exists. Try again: ", term))
	}

	return term
}

func (c *Cards) requestDefinition(num int) string {
	var definition string
	c.print(fmt.Sprintf("The definition for card #%d:", num+1))
	for true {
		definition = strings.TrimSpace(c.getInput())
		if !c.checkDefinitionAlreadyExists(definition) {
			return definition
		}
		c.print(fmt.Sprintf("The definition \"%s\" already exists. Try again:", definition))
	}

	return definition
}

func (c *Cards) requestCardDefinition() string {
	var definition string
	c.print("The definition for the card: ")
	for true {
		definition = strings.TrimSpace(c.getInput())
		if !c.checkDefinitionAlreadyExists(definition) {
			return definition
		}
		c.print(fmt.Sprintf("The definition \"%s\" already exists. Try again: ", definition))
	}

	return definition
}

func (c *Cards) checkTermAlreadyExists(term string) bool {
	for _, card := range c.CardStack {
		if card.Term == term {
			return true
		}
	}

	return false
}

func (c *Cards) checkDefinitionAlreadyExists(definition string) bool {
	for _, card := range c.CardStack {
		if card.Definition == definition {
			return true
		}
	}

	return false
}

func (c *Cards) findTermByDefinition(definition string) (card Card, err error) {
	for _, card := range c.CardStack {
		if card.Definition == definition {
			return card, err
		}
	}

	err = errors.New("definition not found")
	return Card{}, err
}

func (c *Cards) addCard(term string, definition string) bool {
	c.CardStack = append(c.CardStack, Card{Term: term, Definition: definition})

	return c.CardStack[len(c.CardStack)-1].Term == term && c.CardStack[len(c.CardStack)-1].Definition == definition
}

func (c *Cards) deleteCardByTermDefinition(term string, definition string) bool {
	for key, currentCard := range c.CardStack {
		if currentCard.Term == term && currentCard.Definition == definition {
			return c.deleteCard(key)
		}
	}

	return false
}

func (c *Cards) deleteCardByTerm(term string) bool {
	for key, currentCard := range c.CardStack {
		if currentCard.Term == term {
			return c.deleteCard(key)
		}
	}

	return false
}

func (c *Cards) deleteCardByCard(searchedCard Card) bool {
	return c.deleteCardByTerm(searchedCard.Term)
}

func (c *Cards) deleteCard(cardNumber int) bool {
	oldLength := len(c.CardStack)
	c.CardStack = append(c.CardStack[:cardNumber], c.CardStack[cardNumber+1:]...)
	return len(c.CardStack)+1 == oldLength
}

func (c *Card) updateCard(term string, definition string) bool {
	c.Term = term
	c.Definition = definition

	return c.Term == term && c.Definition == definition
}

func (c *Cards) updateCardByOldValues(term string, definition string, oldTerm string, oldDefinition string) bool {
	for _, currentCard := range c.CardStack {
		if currentCard.Term == oldTerm && currentCard.Definition == oldDefinition {
			return currentCard.updateCard(term, definition)
		}
	}

	return false
}

func (c *Cards) updateCard(term string, definition string, cardNumber int) bool {
	c.CardStack[cardNumber].updateCard(term, definition)

	return c.CardStack[cardNumber].Term == term && c.CardStack[cardNumber].Definition == definition
}

func (c *Cards) solveAll() bool {
	scanner := bufio.NewScanner(os.Stdin)
	var advisedSolution string
	var result bool
	result = true
	for _, currentCard := range c.CardStack {
		c.print(fmt.Sprintf("Print the definition of \"%s\"", currentCard.Term))
		scanner.Scan()

		advisedSolution = scanner.Text()

		if currentCard.solve(advisedSolution) {
			c.print("Correct!")
		} else {
			otherCard, err := c.findTermByDefinition(advisedSolution)
			if err == nil {
				c.print(fmt.Sprintf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".",
					currentCard.getDefinition(),
					otherCard.Term))
			} else {
				c.print(fmt.Sprintf("Wrong. The right answer is \"%s\".", currentCard.getDefinition()))
			}
			result = false
		}
	}

	return result
}

func (c *Cards) ask(amount int) {
	for i := 0; i < amount; i++ {
		randomIndex := rand.Intn(len(c.CardStack))

		c.print(fmt.Sprintf("Print the definition of \"%s\": ", c.CardStack[randomIndex].Term))
		input := c.getInput()
		if c.CardStack[randomIndex].solve(input) {
			c.CardStack[randomIndex].Tries++
			c.print("Correct!")
		} else {
			otherCard, err := c.findTermByDefinition(input)
			c.CardStack[randomIndex].Tries++
			c.CardStack[randomIndex].Failures++
			if err == nil {
				c.print(fmt.Sprintf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".",
					c.CardStack[randomIndex].getDefinition(),
					otherCard.Term))
			} else {
				c.print(fmt.Sprintf("Wrong. The right answer is \"%s\".", c.CardStack[randomIndex].getDefinition()))
			}
		}
	}
}

func (c *Cards) solve(term string, definition string) bool {
	for _, currentCard := range c.CardStack {
		if currentCard.Term == term {
			return currentCard.solve(definition)
		}
	}

	return false
}

func (c *Cards) reset() bool {
	result := true
	var newStack []Card
	for _, card := range c.CardStack {
		card.Tries = 0
		card.Failures = 0
		newStack = append(newStack, card)
	}

	c.CardStack = newStack
	for _, card := range c.CardStack {
		if card.Tries != 0 || card.Failures != 0 {
			result = false
		}
	}

	return result
}

func (c *Cards) getHardestCard() (count int, cards []string) {
	var ct int = 0
	var cds []string

	for _, card := range c.CardStack {
		if ct == 0 || ct < card.Failures {
			ct = card.Failures
			cds = []string{card.Term}
		} else if ct == card.Failures {
			cds = append(cds, card.Term)
		}
	}

	return ct, cds
}

func (c *Card) solve(definition string) bool {
	return c.Definition == definition
}

func (c *Card) getDefinition() string {
	return c.Definition
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

	//	c.print(cards[term])
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
	fmt.Println("cardStack:")
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

func IntSliceContains(numbers []int, searchedNumber int) bool {
	result := false
	for _, value := range numbers {
		if value == searchedNumber {
			result = true
			break
		}
	}

	return result
}
