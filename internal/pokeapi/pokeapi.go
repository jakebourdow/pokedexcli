package pokeapi

import (
	"net/http"
	"time"

	"github.com/jakebourdow/pokedexcli/internal/pokecahe"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecahe.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecahe.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
