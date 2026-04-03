package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/michaelptak/pokedex-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeApiClient: pokeClient,
	}

	// Start repl
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())

		commandName := userInput[0]
		args := []string{}
		if len(userInput) > 1 {
			args = userInput[1:]
		}
		command, ok := getCommands()[commandName]

		if ok {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}
