package main

import (
	"fmt"
	"math"
)

func main() {
	// write your code here
	var ingredientsInMachine = map[string]int{
		"water":          400,
		"milk":           540,
		"coffeeBeans":    120,
		"disposableCups": 9,
		"money":          550,
	}
	printStatistics(ingredientsInMachine)
	var running bool = true
	for running {
		switch menu() {
		case "buy":
			ingredientsInMachine = buy(ingredientsInMachine)
			break
		case "fill":
			ingredientsInMachine = fill(ingredientsInMachine)
			break
		case "take":
			ingredientsInMachine = take(ingredientsInMachine)
			break
		case "remaining":
			printStatistics(ingredientsInMachine)
			break
		case "exit":
			running = false
			break
		default:
			printStatistics(ingredientsInMachine)
		}
	}
}

func menu() string {
	var choosedOption string
	fmt.Println("Write action (buy, fill, take):")
	fmt.Scan(choosedOption)

	return choosedOption
}

func buy(ingredientsInMachine map[string]int) map[string]int {
	var coffeeIngredients = map[string]map[string]int{
		"espresso": {
			"water":       250,
			"milk":        0,
			"coffeeBeans": 16,
			"money":       4,
		},
		"latte": {
			"water":       350,
			"milk":        75,
			"coffeeBeans": 20,
			"money":       7,
		},
		"cappuccino": {
			"water":       200,
			"milk":        100,
			"coffeeBeans": 12,
			"money":       6,
		},
	}

	var coffeeType int
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	fmt.Scan(&coffeeType)

	switch coffeeType {
	case 1:
		ingredientsInMachine["milk"] -= coffeeIngredients["espresso"]["milk"]
		ingredientsInMachine["water"] -= coffeeIngredients["espresso"]["water"]
		ingredientsInMachine["coffeeBeans"] -= coffeeIngredients["espresso"]["coffeeBeans"]
		ingredientsInMachine["money"] += coffeeIngredients["espresso"]["money"]
		ingredientsInMachine["disposableCups"] -= 1
		break
	case 2:
		ingredientsInMachine["milk"] -= coffeeIngredients["latte"]["milk"]
		ingredientsInMachine["water"] -= coffeeIngredients["latte"]["water"]
		ingredientsInMachine["coffeeBeans"] -= coffeeIngredients["latte"]["coffeeBeans"]
		ingredientsInMachine["money"] += coffeeIngredients["latte"]["money"]
		ingredientsInMachine["disposableCups"] -= 1
		break
	case 3:
		ingredientsInMachine["milk"] -= coffeeIngredients["cappuccino"]["milk"]
		ingredientsInMachine["water"] -= coffeeIngredients["cappuccino"]["water"]
		ingredientsInMachine["coffeeBeans"] -= coffeeIngredients["cappuccino"]["coffeeBeans"]
		ingredientsInMachine["money"] += coffeeIngredients["cappuccino"]["money"]
		ingredientsInMachine["disposableCups"] -= 1
		break
	}

	return ingredientsInMachine
}

func fill(ingredientsInMachine map[string]int) map[string]int {
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(ingredientsInMachine["water"])
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(ingredientsInMachine["milk"])
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(ingredientsInMachine["coffeeBeans"])
	fmt.Println("Write how many disposable cups of coffee you want to add:")
	fmt.Scan(ingredientsInMachine["disposableCups"])

	return ingredientsInMachine
}

func take(ingredientsInMachine map[string]int) map[string]int {
	fmt.Println(fmt.Sprintf("I gave you $%d", ingredientsInMachine["money"]))
	ingredientsInMachine["money"] = 0
	return ingredientsInMachine
}

func printStatistics(ingredientsInMachine map[string]int) {
	fmt.Println("The coffee machine has:")
	fmt.Println(fmt.Sprintf("%d ml of water", ingredientsInMachine["water"]))
	fmt.Println(fmt.Sprintf("%d ml of milk", ingredientsInMachine["milk"]))
	fmt.Println(fmt.Sprintf("%d g of coffee beans", ingredientsInMachine["coffeeBeans"]))
	fmt.Println(fmt.Sprintf("%d disposable cups", ingredientsInMachine["disposableCups"]))
	fmt.Println(fmt.Sprintf("$%d of moneyr", ingredientsInMachine["money"]))
}

func checkCoffeeMaker() {
	var water, milk, coffeeBeans, coffeeAmount int
	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(water)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(milk)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(coffeeBeans)
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(coffeeAmount)

	if coffeeAmount == 0 {
		fmt.Println("Yes, I can make that amount of coffee ")
		return
	}

	ingredientsInMachine := map[string]int{"water": water, "milk": milk, "coffee_beans": coffeeBeans}
	neededIngredients := calculateCoffeeIngredients(coffeeAmount)
	brewable := brewableCoffeeAmount(ingredientsInMachine, neededIngredients)

	if checkCoffeeBrewable(ingredientsInMachine, neededIngredients) {
		if brewable >= coffeeAmount {
			fmt.Println(fmt.Sprintf("Yes, I can make that amount of coffee (and even %d more than that)",
				brewable-coffeeAmount))
		} else {
			fmt.Println("Yes, I can make that amount of coffee ")
		}
	} else {
		fmt.Println("No, I can make only 0 cups of coffee")
	}
}

func brewableCoffeeAmount(ingredientsInMachine map[string]int, neededIngredients map[string]int) int {
	milkAmount := math.Floor(float64(ingredientsInMachine["milk"] / neededIngredients["milk"]))
	waterAmount := math.Floor(float64(ingredientsInMachine["water"] / neededIngredients["water"]))
	beansAmount := math.Floor(float64(ingredientsInMachine["coffee_beans"] / neededIngredients["coffee_beans"]))

	return int(math.Min(math.Min(milkAmount, waterAmount), beansAmount))
}

func checkCoffeeBrewable(ingredientsInMachine map[string]int, neededIngredients map[string]int) bool {
	return ingredientsInMachine["water"] >= neededIngredients["water"] &&
		ingredientsInMachine["milk"] >= neededIngredients["milk"] &&
		ingredientsInMachine["coffee_beans"] >= neededIngredients["coffee_beans"]
}

func coffeeCalculator() {
	var coffeeAmount int
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&coffeeAmount)
	ingredients := calculateCoffeeIngredients(coffeeAmount)
	fmt.Println(fmt.Sprintf("For %d cups of coffee you will need:", coffeeAmount))
	fmt.Println(fmt.Sprintf("%d ml of water", ingredients["water"]))
	fmt.Println(fmt.Sprintf("%d ml of milk", ingredients["milk"]))
	fmt.Println(fmt.Sprintf("%d g of coffee beans", ingredients["coffee_beans"]))
}

func calculateCoffeeIngredients(numberOfCoffees int) map[string]int {
	coffee := make(map[string]int)
	coffee["water"] = 200 * numberOfCoffees
	coffee["milk"] = 50 * numberOfCoffees
	coffee["coffee_beans"] = 15 * numberOfCoffees

	return coffee
}

func print_text() {
	fmt.Print("Starting to make a coffee\nGrinding coffee beans\nBoiling water\nMixing boiled water with " +
		"crushed coffee beans\nPouring coffee into the cup\nPouring some milk into the cup\nCoffee is ready!")
}
