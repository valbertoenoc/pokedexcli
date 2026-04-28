package main

import (
	"fmt"

	"github.com/valbertoenoc/pokedexcli/internal/pokeapi"
)

func commandMapb(config *Config) error {
	locationURL := pokeapi.LocationURL
	if config.previousURL != nil {
		fmt.Printf("using previousURL %v:", &locationURL)
		locationURL = *config.nextURL
	}

	locationResponse, err := pokeapi.GetLocations(locationURL)
	if err != nil {
		return fmt.Errorf("could not fetch locations %v", err)
	}

	config.previousURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Printf("location-> %s\n", location.Name)
	}

	return nil
}
