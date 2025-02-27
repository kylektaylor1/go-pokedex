package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/kylektaylor1/go-pokedex/internal/pokeapi"
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
	url := config.nextLocationsUrl
	entries := config.pokecache.Entries

	entriesList := make([]string, 0)
	for key := range entries {
		entriesList = append(entriesList, key)
	}
	fmt.Println(entriesList)

	if url != nil {
		fmt.Println("url: ", *url)
		if entry, ok := entries[*url]; ok {
			fmt.Println("hitting cache")
			var pokedata pokeapi.PokeapiData
			unmErr := json.Unmarshal(entry.Val, &pokedata)
			if unmErr != nil {
				return unmErr
			}
			config.nextLocationsUrl = &pokedata.Next
			config.prevLocationsUrl = &pokedata.Previous

			locations := pokedata.Results
			for _, l := range locations {
				fmt.Println(l.Name)
			}
			return nil
		}
	}

	fmt.Println("NOT hitting cache")
	locResp, err := config.pokeapiClient.ListLocations(config.nextLocationsUrl)
	if err != nil {
		return err
	}
	cache := config.pokecache
	bytes, err := json.Marshal(locResp)
	if err != nil {
		return err
	}

	if url != nil {
		err := cache.Add(*url, bytes)
		if err != nil {
			return errors.New("error adding to cache")
		}
	} else {
		err := cache.Add(pokeapi.BaseURL+"/location-area", bytes)
		if err != nil {
			return errors.New("error adding to cache")
		}
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

	url := config.prevLocationsUrl
	entries := config.pokecache.Entries

	fmt.Println("url: ", *url)
	entriesList := make([]string, 0)
	for key := range entries {
		entriesList = append(entriesList, key)
	}
	fmt.Println(entriesList)

	if entry, ok := entries[*url]; ok {
		fmt.Println("hitting cache")
		var pokedata pokeapi.PokeapiData
		unmErr := json.Unmarshal(entry.Val, &pokedata)
		if unmErr != nil {
			return unmErr
		}
		config.nextLocationsUrl = &pokedata.Next
		config.prevLocationsUrl = &pokedata.Previous

		locations := pokedata.Results
		for _, l := range locations {
			fmt.Println(l.Name)
		}
		return nil
	}

	fmt.Println("NOT hitting cache")

	locResp, err := config.pokeapiClient.ListLocations(config.prevLocationsUrl)
	if err != nil {
		return err
	}
	cache := config.pokecache
	bytes, err := json.Marshal(locResp)
	if err != nil {
		return err
	}
	cache.Add(*url, bytes)

	config.nextLocationsUrl = &locResp.Next
	config.prevLocationsUrl = &locResp.Previous

	for _, l := range locResp.Results {
		fmt.Println(l.Name)
	}
	return nil
}
