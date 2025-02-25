package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]CliCommand

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func main() {
	commands = map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()

		switch input {
		case "exit":
			commands["exit"].callback()
		case "help":
			commands["help"].callback()
		default:
			fmt.Println("Unknown command")

		}
	}
}

func cleanInput(input string) []string {
	lower := strings.ToLower(input)
	words := strings.Fields(lower)
	return words
}
