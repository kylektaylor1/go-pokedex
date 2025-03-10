package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func (c *Client) GetPokemon(pageUrl *string, name *string) (PokeapiPokemonResponse, error) {
	if name == nil {
		return PokeapiPokemonResponse{}, nil
	}

	url := BaseURL + "/pokemon" + "/" + *name
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		fmt.Println("hitting cache")
		var pokedata PokeapiPokemonResponse
		unmErr := json.Unmarshal(val, &pokedata)
		if unmErr != nil {
			return PokeapiPokemonResponse{}, unmErr
		}

		return pokedata, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeapiPokemonResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeapiPokemonResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeapiPokemonResponse{}, err
	}

	var pokedata PokeapiPokemonResponse
	unmErr := json.Unmarshal(body, &pokedata)
	if unmErr != nil {
		return PokeapiPokemonResponse{}, err
	}

	c.cache.Add(url, body)

	return pokedata, nil
}

func AttemptCatchPokemon(exp int) bool {
	fmt.Printf("exp: %v\n", exp)
	val := rand.Intn(exp / 10)
	userVal := rand.Intn(exp / 10)
	fmt.Println(val, userVal)

	return val == userVal
}

func (c *Client) CatchPokemon(name string, pokemon PokeapiPokemonResponse) error {
	if _, ok := c.pokedex[name]; ok {
		return errors.New("pokemon already caught")
	}

	c.pokedex[name] = pokemon

	return nil
}

func (c *Client) InspectPokemon(name string) error {
	if val, ok := c.pokedex[name]; ok {
		var hp int
		var attack int
		var defense int
		var specialAttack int
		var specialDefense int
		var speed int
		for _, s := range val.Stats {
			if s.Stat.Name == "hp" {
				hp = s.BaseStat
			}
			if s.Stat.Name == "attack" {
				attack = s.BaseStat
			}
			if s.Stat.Name == "defense" {
				defense = s.BaseStat
			}
			if s.Stat.Name == "special-attack" {
				specialAttack = s.BaseStat
			}
			if s.Stat.Name == "special-defense" {
				specialDefense = s.BaseStat
			}
			if s.Stat.Name == "speed" {
				speed = s.BaseStat
			}
		}
		fmt.Println("Name: ", val.Name)
		fmt.Println("Height: ", val.Height)
		fmt.Println("Weight: ", val.Weight)
		fmt.Println("Stats:")
		fmt.Println("  -hp: ", hp)
		fmt.Println("  -attack: ", attack)
		fmt.Println("  -defense: ", defense)
		fmt.Println("  -special-attack: ", specialAttack)
		fmt.Println("  -special-defense: ", specialDefense)
		fmt.Println("  -speed: ", speed)
		fmt.Println("Types:", speed)
		for _, t := range val.Types {
			fmt.Println("- ", t.Type.Name)
		}
		return nil
	}

	return errors.New("you have not caught that pokemon")
}

func (c *Client) InspectPokedex() error {
	fmt.Println("Your Pokedex:")
	for _, p := range c.pokedex {
		fmt.Println(p.Name)
	}
	return nil
}
