package main

import "fmt"

func commandHelp(*config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for k, v := range getCommands() {
		fmt.Printf("%v: %v\n", k, v.description)
	}
	fmt.Println()
	return nil
}
