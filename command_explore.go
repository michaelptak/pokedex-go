package main

import (
	"errors"
	"fmt"
)

func commandExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location name. Use `map` to find some.")
	}

	name := args[0]
	location, err := conf.pokeApiClient.ListPokemon(name)
	if err != nil {
		return nil
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")

	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil

}
