package pokeapi

import (
	"net/http"
	"time"

	"github.com/marcel-alter/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	httpClient  http.Client
	ClientCache pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	newCache := pokecache.NewCache(cacheInterval)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		ClientCache: newCache,
	}
}
