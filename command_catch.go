package main

import (
	"errors"
	"fmt"
)

func commandCatch(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a Pokemon name.")
	}

	name := args[0]
	location, err := conf.pokeApiClient.ListPokemon(name)
	if err != nil {
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")

	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil

}
