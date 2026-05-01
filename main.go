package main

import (
	"time"

	"github.com/valbertoenoc/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Second*5)
	cfg := &config{
		pokeapiClient:   pokeClient,
		capturedPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
