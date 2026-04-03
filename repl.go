package main

import (
	"strings"

	"github.com/michaelptak/pokedex-go/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowerStr := strings.ToLower(text)
	fields := strings.Fields(lowerStr)
	return fields
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeApiClient pokeapi.Client
	Next          *string
	Previous      *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Get command info",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List page of 20 pokemon locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back to previous 20 locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
	}

}
