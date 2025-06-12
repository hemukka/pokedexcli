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

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)
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

	var locationsData LocationAreas
	if err = json.Unmarshal(data, &locationsData); err != nil {
		return LocationAreas{}, err
	}

	// add the data to cache
	c.cache.Add(url, data)

	return locationsData, nil
}

// structure of response from the paginated resource list from "Location Areas" endpoint
type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationArea(area string) (LocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + area

	if data, ok := c.cache.Get(url); ok {
		var locationData LocationArea
		if err := json.Unmarshal(data, &locationData); err != nil {
			return LocationArea{}, err
		}
		fmt.Println("(data from cache)")
		return locationData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		// 404 probably means that there is no area with the given name
		return LocationArea{}, fmt.Errorf(
			"location area doesn't exist",
		)
	}
	if res.StatusCode != 200 {
		return LocationArea{}, fmt.Errorf(
			"response failed with status code: %v",
			res.StatusCode,
		)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var locationsData LocationArea
	if err = json.Unmarshal(data, &locationsData); err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, data)

	return locationsData, nil
}

// structure of response from "Location Areas" endpoint when id or name is given
type LocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
