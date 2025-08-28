package main

import (
  "fmt"
	"errors"
)

func commandCatch(cfg *config, args ...string) error {
 	if len(args) != 1 {
		return errors.New("You need to provide which pokemon you want to catch")
	} 

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.PokemonStats(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

  return nil
}
