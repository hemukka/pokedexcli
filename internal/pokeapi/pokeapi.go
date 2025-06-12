package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocationAreas(url string) (locationAreas, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	res, err := http.Get(url)
	if err != nil {
		return locationAreas{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return locationAreas{}, fmt.Errorf(
			"response failed with status code: %v",
			res.StatusCode,
		)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreas{}, err
	}

	var locationsData locationAreas
	if err = json.Unmarshal(data, &locationsData); err != nil {
		return locationAreas{}, err
	}

	return locationsData, nil
}

type locationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
