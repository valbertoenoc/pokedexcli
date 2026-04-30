package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a location name")
	}

	area := string(args[0])
	fmt.Printf("Exploring %s\n", area)
	pokemonList, err := cfg.pokeapiClient.GetLocationArea(area)
	if err != nil {
		return fmt.Errorf("error fetching pokemon list: %v", err)
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range pokemonList.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
