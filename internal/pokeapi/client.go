package pokeapi

import (
	"github.com/lalobec/pokedexGo/internal/pokecache"
)

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient  http.Client
	cacheClient *pokecache.Cache
}

// If the API takes more time than timeout due to any
// reason, the request will be cancelled to prevent the
// CLI from hanging indefinitely
func NewClient(timeout time.Duration, interval time.Duration) Client {
	cacheC := pokecache.NewCache(interval)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cacheClient: cacheC,
	}
}
