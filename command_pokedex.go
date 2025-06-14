package main

import (
	"errors"
	"fmt"
)

func commandPokedex(config *config, _ string) error {
	if len(config.pokedex) == 0 {
		return errors.New("you haven't caught any pokemon")
	}

	println("Your Pokedex:")
	for _, pokemon := range config.pokedex {
		fmt.Println(" - " + pokemon.Name)
	}
	return nil
}
