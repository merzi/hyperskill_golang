package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 	write your code here
	var running = true
	var reader = bufio.NewReader(os.Stdin)
	var notes []string
	var maxNoteNumber int

	for {
		fmt.Print("Enter the maximum number of notes:")
		_, err := fmt.Scanf("%d", &maxNoteNumber)

		if err == nil {
			break
		}

		_, err = reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println("Sorry, invalid input. Please enter an integer: ")
	}

	for running {
		fmt.Print("Enter a command and data: ")
		str, err := reader.ReadString('\n')

		if err != nil {
			continue
		}

		if checkNoteCommand(str, &notes) == false {
			continue
		}

		switch strings.TrimSpace(strings.Split(str, " ")[0]) {
		case "exit":
			running = false
			fmt.Println("[Info] Bye!")
			break
		case "create":
			create(strings.TrimSpace(strings.Replace(str, "create", "", 1)), &notes, &maxNoteNumber)
			break
		case "list":
			list(&notes)
			break
		case "clear":
			clear(&notes)
			break
		case "update":
			intVar, _ := strconv.Atoi(strings.Split(str, " ")[1])
			str = strings.Replace(str, "update", "", 1)
			str = strings.Replace(str, string(rune(intVar)), "", 1)
			updateNote(intVar, strings.TrimSpace(str), &notes)
			break
		case "delete":
			intVar, _ := strconv.Atoi(strings.Split(str, " ")[1])
			deleteNote(intVar, &notes)
			break
		default:
			fmt.Print("[Error] Unknown command\n")
		}
	}
}

func checkNoteCommand(command string, notes *[]string) bool {
	var commands = []string{"create", "update", "delete"}
	if containsStringFromSlice(command, commands) == false {
		return true
	} else if len(strings.TrimSpace(replaceStringFromSlice(command, commands, ""))) < 1 {
		fmt.Print("[Error] Missing note argument\n")
		return false
	} else if strings.Contains(command, "create") && len(strings.Split(command, " ")) >= 2 {
		return true
	} else if containsStringFromSlice(command, []string{"update", "delete"}) {
		intVar, err := strconv.Atoi(strings.Split(command, " ")[1])
		if err != nil {
			fmt.Print(fmt.Sprintf("[Error] Invalid position: %s\n", strings.Split(command, " ")[1]))
			fmt.Print("[Error] Missing position argument\n")
			return false
		} else if len(*notes) <= intVar {
			fmt.Print(fmt.Sprintf("[Error] Position %d is out of the boundaries [1, %d]\n", intVar, len(*notes)))
			return false
		}

	}

	fmt.Print("[Error] Missing note argument\n")
	return false
}

func containsStringFromSlice(s string, stringSlice []string) bool {
	for _, value := range stringSlice {
		if strings.Contains(s, value) {
			return true
		}
	}

	return false
}

func replaceStringFromSlice(s string, stringSlice []string, replacement string) string {
	for _, value := range stringSlice {
		s = strings.Replace(s, value, replacement, 1)
	}

	return s
}

func create(data string, notes *[]string, maxNoteNumber *int) bool {
	if len(*notes) < *maxNoteNumber {
		n := *notes
		*notes = append(n, data)
		fmt.Print("[OK] The note was successfully created\n")
		return true
	}

	fmt.Print("[Error] Notepad is full\n")
	return false
}

func clear(notes *[]string) bool {
	*notes = []string{}
	if len(*notes) > 0 {
		fmt.Print("[Error] deletion failed!")
		return false
	}

	fmt.Print("[OK] All notes were successfully deleted\n")
	return true
}

func list(notes *[]string) {
	if len(*notes) < 1 {
		fmt.Print("[Info] Notepad is empty\n")
		return
	}

	for key, value := range *notes {
		fmt.Print(fmt.Sprintf("[Info] %d: %s\n", key+1, value))
	}
}

func updateNote(number int, newNote string, notes *[]string) bool {
	number -= 1
	if len(*notes) >= number {
		b := *notes
		b[number] = newNote
		*notes = b
		fmt.Print(fmt.Sprintf("[OK] The note at position %d was successfully updated\n", number))
		return true
	}
	fmt.Print(fmt.Sprintf("[Error] update of index %d failed!", number+1))
	return false
}

func deleteNote(number int, notes *[]string) bool {
	number -= 1
	if len(*notes) >= number {
		var oldCount = len(*notes)
		b := *notes
		*notes = removeFromStringSlice(b, number)

		if len(*notes) == oldCount-1 {
			fmt.Print(fmt.Sprintf("[OK] The note at position %d was successfully deleted\n", number+1))
			return true
		}
	}
	fmt.Print(fmt.Sprintf("[Error] deletion of index %d failed!", number+1))
	return false
}

func removeFromStringSlice(slice []string, num int) []string {
	return append(slice[:num], slice[num+1:]...)
}
