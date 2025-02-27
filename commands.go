package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	commands := getCommands()
	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandMapf(config *Config) error {
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

func commandMapb(config *Config) error {
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
