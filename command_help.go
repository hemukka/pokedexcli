package main

import "fmt"

func commandHelp(*config, string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, v := range getCommands() {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	fmt.Println()
	return nil
}
