package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/valbertoenoc/pokedexcli/internal/pokeapi"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cliCommand, ok := getCommands()[commandName]
		if !ok {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}

		err := cliCommand.callback(cfg, args...)
		if err != nil {
			fmt.Printf("Error executing command %s: %v\n", commandName, err)
		}
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(strings.ToLower(text))
	return strings.Fields(trimmed)
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

type config struct {
	pokeapiClient   pokeapi.Client
	capturedPokemon map[string]pokeapi.Pokemon
	nextURL         *string
	previousURL     *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"mapf": {
			name:        "mapf",
			description: "Fetches batch of 20 next location areas.",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Fetches batch of 20 previous location areas.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Fetches list of pokemon present in the area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catche Pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Display detailed Pokemon data.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display all captured pokemon.",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
	}
}
