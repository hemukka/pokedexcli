package main

import (
	"fmt"
	"os"
)

func commandExit(*config, string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
