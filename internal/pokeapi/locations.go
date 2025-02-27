package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (PokeapiData, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
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
	return pokedata, nil
}
