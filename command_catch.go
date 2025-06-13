package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *config, name string) error {
	if name == "" {
		return fmt.Errorf("argument missing: pokemon name")
	}

	pokemon, err := config.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	// max base_experience is 255
	power := rand.Intn(256)
	if power >= pokemon.BaseExperience {
		fmt.Printf("%v was caught!\n", name)
		config.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", name)
	}

	return nil
}
