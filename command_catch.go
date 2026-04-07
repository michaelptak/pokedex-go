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

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	baseXP := pokemon.BaseExperience
	catch_chance := rand.Intn(baseXP)
	if catch_chance < 100 {
		fmt.Printf("Successfully added %s to your Pokedex!\n", pokemon.Name)
		conf.Pokedex[name] = pokemon
	} else {
		fmt.Printf("You failed to catch the pokemon!\n")
	}

	return nil

}
