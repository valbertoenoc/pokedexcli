package main

import (
	"time"

	"github.com/valbertoenoc/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second * 5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
