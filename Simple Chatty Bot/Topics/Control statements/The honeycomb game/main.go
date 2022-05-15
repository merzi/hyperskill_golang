package main

import "fmt"

func main() {
	var emoji = "❓"
	fmt.Scanf("%s", &emoji)

	// Please do not delete the emojis after the case statement, just fix the code errors.
	// Also please do not delete or change the text within the fmt.Println functions!
	switch emoji {
	case "⭕":
		fmt.Println("You have picked the circle. Not the easiest shape!")
		break
	case "🔺":
		fmt.Println("You have picked the triangle. The easiest shape!")
		break
	case "⭐":
		fmt.Println("You have picked the star. Easier than circle, harder than triangle.")
		break
	case "☂️":
		fmt.Println("You have picked the umbrella. This is the hardest shape! GOOD LUCK.")
		break
	default:
		fmt.Println("You have picked an invalid emoji. Please try again or be eliminated from the game.")
	}
}
