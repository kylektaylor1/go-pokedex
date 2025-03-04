package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationAreaByName(pageUrl *string) (PokeapiLocationAreaResp, error) {
	url := BaseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		fmt.Println("hitting cache")
		var pokedata PokeapiLocationAreaResp
		unmErr := json.Unmarshal(val, &pokedata)
		if unmErr != nil {
			return PokeapiLocationAreaResp{}, unmErr
		}

		return pokedata, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeapiLocationAreaResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeapiLocationAreaResp{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeapiLocationAreaResp{}, err
	}

	var pokedata PokeapiLocationAreaResp
	unmErr := json.Unmarshal(body, &pokedata)
	if unmErr != nil {
		return PokeapiLocationAreaResp{}, err
	}

	c.cache.Add(url, body)

	return pokedata, nil
}
