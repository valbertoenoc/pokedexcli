package main

import (
	"fmt"
	"os"
)

func commandExit(config *config) error {
	fmt.Println()
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
