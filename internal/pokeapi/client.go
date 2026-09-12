package pokeapi

import (
	"net/http"
	"time"

	"github.com/dyn64/pokedexcli/internal/pokecache"
)

// Struct for the httpClient with cache to keep requests for x minutes
type Client struct {
	httpClient http.Client
	pokeCache  pokecache.Cache
}

// Creates a new Client with a set timeout + cache with a set interval for how long the cache stores records
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: pokecache.NewCache(cacheInterval),
	}
}
