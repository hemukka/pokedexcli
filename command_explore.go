package main

import "fmt"

func commandExplore(config *config, param string) error {
	name := param
	if name == "" {
		return fmt.Errorf("argument missing: location area name")
	}

	fmt.Printf("Exploring %v...\n", name)
	data, err := config.pokeapiClient.GetLocationArea(name)
	if err != nil {
		return err
	}

	if len(data.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon found in this area")
		return nil
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range data.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
