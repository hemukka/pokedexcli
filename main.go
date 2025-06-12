package main

import (
	"time"

	"github.com/hemukka/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(
			5*time.Second,
			5*time.Minute,
		),
	}
	repl(cfg)
}
