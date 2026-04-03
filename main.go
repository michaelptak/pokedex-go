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
