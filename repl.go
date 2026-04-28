package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	config := &Config{}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		cliCommand, ok := getCommands()[commandName]
		if !ok {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}

		err := cliCommand.callback(config)
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
	callback    func(*Config) error
}

type Config struct {
	nextURL     *string
	previousURL *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Fetches batch of 20 next location areas.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Fetches batch of 20 previous location areas.",
			callback:    commandMapb,
		},
	}
}
