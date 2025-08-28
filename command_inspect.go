package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You need to provide a pokemon to see the details")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return errors.New("This pokemon does not exist in your pokedex")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("---------------------------------")
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %v: %v \n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("---------------------------------")
	fmt.Printf("Types:\n")
	for _, poketype := range pokemon.Types {
		fmt.Printf("  - %v\n", poketype.Type.Name)
	}

	return nil

}
