package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Config struct {
	Next     string
	Previous *string
}

type CliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type PokeapiData struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []PokeLocation `json:"results"`
}

type PokeLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getCommands(config *Config) map[string]CliCommand {
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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get location areas - b",
			callback:    commandMapb,
		},
	}
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	commands := getCommands(&Config{Next: "", Previous: nil})
	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func getPokeLocations(url string, config *Config) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var pokedata PokeapiData
	unmErr := json.Unmarshal(body, &pokedata)
	if unmErr != nil {
		return nil, err
	}
	results := pokedata.Results
	// set next
	config.Next = pokedata.Next
	config.Previous = &pokedata.Previous

	locations := make([]string, 0)
	for _, r := range results {
		locations = append(locations, r.Name)
	}

	return locations, nil
}

func commandMap(config *Config) error {
	url := config.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	locations, err := getPokeLocations(url, config)
	if err != nil {
		return err
	}
	for _, l := range locations {
		fmt.Println(l)
	}

	return nil
}

func commandMapb(config *Config) error {
	url := *config.Previous
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	locations, err := getPokeLocations(url, config)
	if err != nil {
		return err
	}
	for _, l := range locations {
		fmt.Println(l)
	}
	return nil
}
