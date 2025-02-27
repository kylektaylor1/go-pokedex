package pokeapi

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
