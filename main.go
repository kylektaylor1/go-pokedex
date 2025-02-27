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
	client := pokeapi.NewClient(time.Second * 5)
	config := &Config{
		pokeapiClient: client,
	}
	commands := getCommands()
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
