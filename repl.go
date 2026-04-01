package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerStr := strings.ToLower(text)
	fields := strings.Fields(lowerStr)
	return fields
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
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Get command info",
			callback:    commandHelp,
		},
	}

}
