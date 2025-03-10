package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kylektaylor1/go-pokedex/internal/pokeapi"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	client := pokeapi.NewClient(time.Second*5, time.Second*5)
	config := &Config{
		pokeapiClient: client,
	}
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()

		cleaned := cleanInput(input)

		var args []string
		if len(cleaned) > 1 {
			args = cleaned[1:]
		} else {
			args = nil
		}

		switch cleaned[0] {
		case "exit":
			commands["exit"].callback(config, args...)
		case "help":
			commands["help"].callback(config, args...)
		case "map":
			commands["map"].callback(config, args...)
		case "mapb":
			commands["mapb"].callback(config, args...)
		case "explore":
			commands["explore"].callback(config, args...)
		case "catch":
			commands["catch"].callback(config, args...)
		case "inspect":
			commands["inspect"].callback(config, args...)
		case "pokedex":
			commands["pokedex"].callback(config, args...)
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
