package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (PokeapiData, error) {
	url := BaseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		fmt.Println("hitting cache")
		var pokedata PokeapiData
		unmErr := json.Unmarshal(val, &pokedata)
		if unmErr != nil {
			return PokeapiData{}, unmErr
		}

		return pokedata, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeapiData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeapiData{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeapiData{}, err
	}

	var pokedata PokeapiData
	unmErr := json.Unmarshal(body, &pokedata)
	if unmErr != nil {
		return PokeapiData{}, err
	}

	c.cache.Add(url, body)

	return pokedata, nil
}
