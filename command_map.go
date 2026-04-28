package main

import (
	"fmt"

	"github.com/valbertoenoc/pokedexcli/internal/pokeapi"
)

func commandMap(config *Config) error {
	locationURL := pokeapi.LocationURL
	if config.nextURL != nil {
		fmt.Printf("using nextURL %v:", &locationURL)
		locationURL = *config.nextURL
	}

	locationResponse, err := pokeapi.GetLocations(locationURL)
	if err != nil {
		return fmt.Errorf("could not fetch locations %v", err)
	}

	config.nextURL = &locationResponse.Next

	for _, location := range locationResponse.Results {
		fmt.Printf("location-> %s\n", location.Name)
	}

	return nil
}
