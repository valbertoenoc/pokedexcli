package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationResponse, err := cfg.pokeapiClient.GetLocations(cfg.nextURL)
	if err != nil {
		return fmt.Errorf("could not fetch locations %v", err)
	}

	cfg.nextURL = locationResponse.Next
	cfg.previousURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Printf("location-> %s\n", location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousURL == nil {
		return errors.New("You're on the first page.")
	}
	locationResponse, err := cfg.pokeapiClient.GetLocations(cfg.previousURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationResponse.Next
	cfg.previousURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Printf("location-> %s\n", location.Name)
	}

	return nil
}
