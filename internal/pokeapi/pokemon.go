package pokeapi

import (
	"encoding/json"
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
