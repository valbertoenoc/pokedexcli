package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

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

		err := cliCommand.callback()
		if err != nil {
			fmt.Printf("Error executing command: %s", commandName)
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
	callback    func() error
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
	}
}
