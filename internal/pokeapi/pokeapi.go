package pokeapi

import (
	"net/http"
	"time"
	"github.com/Lasavlad/pokedox/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheinterval time.Duration) Client {
	return Client {
		cache: pokecache.NewCache(cacheinterval),
		httpClient: http.Client{ 
			Timeout: time.Minute,
		},

	}
}

