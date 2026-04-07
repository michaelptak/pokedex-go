package main

import (
	"errors"
	"fmt"
)

func commandInspect(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a Pokemon name.")
	}

	name := args[0]
	pokemon, ok := conf.Pokedex[name]

	if !ok {
		return errors.New("You have not caught that pokemon.")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	return nil
}
