package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
)

const (
	ceilingExperience = 400
	maxChance         = 0.85
	minChance         = 0.15
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return errors.New("error fetching pokemon")
	}

	chance := 1 - float64(pokemon.BaseExperience)/ceilingExperience
	chance = math.Max(float64(chance), minChance)
	chance = math.Min(float64(chance), maxChance)

	roll := rand.Float64()
	if roll < chance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.capturedPokemon[pokemon.Name] = pokemon

	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
