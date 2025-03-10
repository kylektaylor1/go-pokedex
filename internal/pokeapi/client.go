package pokeapi

import (
	"net/http"
	"time"

	"github.com/kylektaylor1/go-pokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	pokedex    Pokedex
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache:      pokecache.NewCache(cacheInterval),
		pokedex:    make(Pokedex),
		httpClient: http.Client{Timeout: timeout},
	}
}
