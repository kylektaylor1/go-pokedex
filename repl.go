package main

import "github.com/kylektaylor1/go-pokedex/internal/pokeapi"

type Config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsUrl *string
	prevLocationsUrl *string
}

type CliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
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
		"map": {
			name:        "map",
			description: "Get location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get location areas - b",
			callback:    commandMapb,
		},
	}
}
