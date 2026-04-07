package main

import "fmt"

func commandPokedex(conf *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	for _, pokemon := range conf.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
