package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *config, name string) error {
	if name == "" {
		return errors.New("argument missing: pokemon name")
	}

	pokemon, caught := config.pokedex[name]
	if !caught {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v dm\n", pokemon.Height)
	fmt.Printf("Weight: %v hg\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, item := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", item.Stat.Name, item.BaseStat)
	}
	fmt.Println("Types:")
	for _, item := range pokemon.Types {
		fmt.Printf("  - %v\n", item.Type.Name)
	}

	return nil
}
