package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/kylektaylor1/go-pokedex/internal/pokeapi"
)

func commandExit(config *Config, params ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, param ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	commands := getCommands()
	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandMapf(config *Config, param ...string) error {
	locResp, err := config.pokeapiClient.ListLocations(config.nextLocationsUrl)
	if err != nil {
		return err
	}

	config.nextLocationsUrl = &locResp.Next
	config.prevLocationsUrl = &locResp.Previous

	locations := locResp.Results
	for _, l := range locations {
		fmt.Println(l.Name)
	}

	return nil
}

func commandMapb(config *Config, param ...string) error {
	if config.prevLocationsUrl == nil {
		return errors.New("you're on the first page")
	}

	locResp, err := config.pokeapiClient.ListLocations(config.prevLocationsUrl)
	if err != nil {
		return err
	}

	config.nextLocationsUrl = &locResp.Next
	config.prevLocationsUrl = &locResp.Previous

	for _, l := range locResp.Results {
		fmt.Println(l.Name)
	}
	return nil
}

func commandExpore(config *Config, params ...string) error {
	if len(params) != 1 {
		return errors.New("you must provide a location area")
	}

	location := params[0]
	fmt.Printf("Exploring %v...\n", location)
	url := pokeapi.BaseURL + "/location-area" + "/" + location

	locResp, err := config.pokeapiClient.LocationAreaByName(&url)
	if err != nil {
		return err
	}

	encounter := locResp.PokemonEncounters
	fmt.Println("Found Pokemon:")
	for _, e := range encounter {
		fmt.Println("- ", e.Pokemon.Name)
	}
	return nil
}

func commandCatch(config *Config, params ...string) error {
	if len(params) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemon := params[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)

	url := pokeapi.BaseURL + "/pokemon" + "/" + pokemon

	data, err := config.pokeapiClient.GetPokemon(&url, &pokemon)
	if err != nil {
		return err
	}

	isCaught := pokeapi.AttemptCatchPokemon(data.BaseExperience)

	if isCaught {
		fmt.Printf("%v was caught\n", pokemon)
	} else {
		fmt.Printf("%v escaped\n", pokemon)
	}

	return nil
}
