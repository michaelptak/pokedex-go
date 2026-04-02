package main

import (
	"bufio"
	"fmt"
	"github.com/michaelptak/pokedex-go/internal/pokeapi"
	"os"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeApiClient: pokeClient,
	}

	// Start repl
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		user_input := cleanInput(scanner.Text())

		commandName := user_input[0]
		command, ok := getCommands()[commandName]

		if ok {
			err := command.callback(cfg)
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
