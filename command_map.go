package main

import (
	"fmt"
)

func commandMap(config *config, _ string) error {
	resp, err := config.pokeapiClient.GetLocationAreas(config.nextURL)
	if err != nil {
		return err
	}

	config.nextURL = resp.Next
	config.prevURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(config *config, _ string) error {
	if config.prevURL == "" {
		return fmt.Errorf("you're on the first page")
	}

	resp, err := config.pokeapiClient.GetLocationAreas(config.prevURL)
	if err != nil {
		return err
	}

	config.nextURL = resp.Next
	config.prevURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
