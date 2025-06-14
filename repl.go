package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hemukka/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       string
	prevURL       string
	pokedex       map[string]pokeapi.Pokemon
}

func repl(config *config) {
	commands := getCommands()

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		argument := ""
		if len(words) >= 2 {
			argument = words[1]
		}

		if command, ok := commands[commandName]; ok {
			err := command.callback(config, argument)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unkown command")
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: " List next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "List Pokemons found in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "   Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: " Print stats of a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught pokemon",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
