package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	prev := ""
	config := &Config{
		Next:     "",
		Previous: &prev,
	}
	commands := getCommands(config)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()
		switch input {
		case "exit":
			commands["exit"].callback(config)
		case "help":
			commands["help"].callback(config)
		case "map":
			commands["map"].callback(config)
		case "mapb":
			commands["mapb"].callback(config)
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
