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

	var reader = bufio.NewReader(os.Stdin)
	var maxNoteNumber int

	maxNoteNumber = getNoteNumber(reader)

	if maxNoteNumber == -1 {
		return
	}

	runner(reader, maxNoteNumber)
}

func getNoteNumber(reader *bufio.Reader) int {
	var maxNoteNumber int
	for {
		fmt.Print("Enter the maximum number of notes: ")
		_, err := fmt.Scanf("%d", &maxNoteNumber)

		if err == nil {
			break
		}

		_, err = reader.ReadString('\n')
		if err != nil {
			return -1
		}
		fmt.Println("Sorry, invalid input. Please enter an integer: ")
	}

	return maxNoteNumber
}

func runner(reader *bufio.Reader, maxNoteNumber int) {
	var running = true
	var notes []string

	for running {
		if menu(&notes, reader, maxNoteNumber) == false {
			running = false
		}
	}
}

func menu(notes *[]string, reader *bufio.Reader, maxNoteNumber int) bool {
	fmt.Print("Enter a command and data: ")
	str, err := reader.ReadString('\n')

	if err != nil {
		return true
	}

	str = strings.Replace(str, "\n", "", -1)
	if checkNoteCommand(str, notes, maxNoteNumber) == false {
		return true
	}

	switch strings.TrimSpace(strings.Split(str, " ")[0]) {
	case "exit":
		fmt.Println("[Info] Bye!")
		return false
		break
	case "create":
		create(strings.TrimSpace(strings.Replace(str, "create", "", 1)), notes, &maxNoteNumber)
		break
	case "list":
		list(notes)
		break
	case "clear":
		clear(notes)
		break
	case "update":
		intVar, err := strconv.Atoi(strings.Split(str, " ")[1])
		if err != nil {
			fmt.Printf("[Error] Invalid position: %s\n", strings.Split(str, " ")[1])
			break
		}

		if len(strings.TrimSpace((*notes)[intVar-1])) < 1 {
			fmt.Println("[Error] There is nothing to update")
			break
		}

		str = strings.TrimSpace(strings.Replace(str, "update", "", 1))
		str = strings.TrimSpace(strings.Replace(str, strings.Split(str, " ")[0], "", 1))
		updateNote(intVar-1, strings.TrimSpace(str), notes)
		break
	case "delete":
		intVar, err := strconv.Atoi(strings.Split(str, " ")[1])
		if err != nil {
			fmt.Printf("[Error] Invalid position: %s\n", strings.Split(str, " ")[1])
			fmt.Print(err)
			break
		}
		deleteNote(intVar-1, notes)
		break
	default:
		fmt.Println("[Error] Unknown command")
	}

	return true
}

func checkNoteCommand(command string, notes *[]string, maxNoteNumbers int) bool {
	var commands = []string{"create", "update", "delete"}
	if containsStringFromSlice(command, commands) == false {
		return true
	} else if checkCommandLength(command, &commands) == false {
		return false
	} else if strings.Contains(command, "create") && len(strings.Split(command, " ")) >= 2 {
		return true
	} else if containsStringFromSlice(command, []string{"update", "delete"}) {
		intVar, err := strconv.Atoi(strings.Split(strings.Replace(command, "\n", "", -1), " ")[1])
		if err != nil {
			fmt.Printf("[Error] Invalid position: %s. Missing position argument\n", strings.Split(command, " ")[1])
			return false
		} else if maxNoteNumbers <= intVar-1 {
			fmt.Printf("[Error] Position %d is out of the boundary [1, %d]\n", intVar, maxNoteNumbers)
			return false
		} else if intVar > len(*notes) ||
			intVar <= maxNoteNumbers && (len(*notes) < intVar-1 || len(strings.TrimSpace((*notes)[intVar-1])) < 1) {
			fmt.Printf("[Error] There is nothing to %s\n", strings.Split(command, " ")[0])
			return false
		} else if intVar <= len(*notes) {
			return true
		}

	}

	fmt.Println("[Error] Missing note argument!")
	return false
}

func checkCommandLength(command string, commandSlice *[]string) bool {
	order := strings.Split(command, " ")[0]

	if len(strings.Split(strings.TrimSpace(command), " ")) < 2 && (order == "update" || order == "delete") {
		fmt.Println("[Error] Missing position argument")
		return false
	} else if len(strings.Split(command, " ")) < 3 && order == "update" {
		fmt.Println("[Error] Missing note argument")
		return false
	} else if len(strings.TrimSpace(replaceStringFromSlice(command, *commandSlice, ""))) < 1 {
		fmt.Println("[Error] Missing note argument")
		return false
	}

	return true
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
		fmt.Println("[OK] The note was successfully created")
		return true
	}

	fmt.Println("[Error] Notepad is full")
	return false
}

func clear(notes *[]string) bool {
	*notes = []string{}
	if len(*notes) > 0 {
		fmt.Print("[Error] deletion failed!")
		return false
	}

	fmt.Println("[OK] All notes were successfully deleted")
	return true
}

func list(notes *[]string) {
	if len(*notes) < 1 {
		fmt.Println("[Info] Notepad is empty")
		return
	}

	for key, value := range *notes {
		fmt.Printf("[Info] %d: %s\n", key+1, value)
	}
}

func updateNote(number int, newNote string, notes *[]string) bool {
	if len(*notes) < number {
		fmt.Println("[Error] There is nothing to update")
		return false
	}

	if len(*notes) > number {
		(*notes)[number] = newNote
		fmt.Printf("[OK] The note at position %d was successfully updated\n", number+1)
		return true
	}
	fmt.Printf("[Error] update of index %d failed!\n", number+1)
	return false
}

func deleteNote(number int, notes *[]string) bool {
	if len(*notes) < number {
		fmt.Println("Error] There is nothing to delete")
		return false
	}

	if len(*notes) >= number {
		var oldCount = len(*notes)
		b := *notes
		*notes = removeFromStringSlice(b, number)

		if len(*notes) == oldCount-1 {
			fmt.Printf("[OK] The note at position %d was successfully deleted\n", number+1)
			return true
		}
	}
	fmt.Print(fmt.Sprintf("[Error] deletion of index %d failed!", number+1))
	return false
}

func removeFromStringSlice(slice []string, num int) []string {
	return append(slice[:num], slice[num+1:]...)
}
