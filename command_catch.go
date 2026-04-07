package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a Pokemon name.")
	}

	name := args[0]
	pokemon, err := conf.pokeApiClient.PokemonInfo(name)
	if err != nil {
		return nil
	}
	catchRoll := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if catchRoll > 50 {
		fmt.Printf("You failed to catch the pokemon!\n")
		return nil
	}

	fmt.Printf("Successfully added %s to your Pokedex!\n", pokemon.Name)
	conf.Pokedex[name] = pokemon

	return nil

}
