package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hemukka/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetLocationAreas(url string) (LocationAreas, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}

	// check if the data is in cache
	if data, ok := c.cache.Get(url); ok {
		var locationsData LocationAreas
		if err := json.Unmarshal(data, &locationsData); err != nil {
			return LocationAreas{}, err
		}
		fmt.Println("(data from cache)")
		return locationsData, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LocationAreas{}, fmt.Errorf(
			"response failed with status code: %v",
			res.StatusCode,
		)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	// add the data to cache
	c.cache.Add(url, data)

	var locationsData LocationAreas
	if err = json.Unmarshal(data, &locationsData); err != nil {
		return LocationAreas{}, err
	}

	return locationsData, nil
}

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
