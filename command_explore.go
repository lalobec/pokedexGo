package main

import (
	"errors"
  "fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location to explore")
	}

	location_name := args[0]
  pokeList, err := cfg.pokeapiClient.PokemonList(location_name) 
  if err != nil {
    return err
  }

  for _, pokemon := range pokeList {
    fmt.Println(pokemon)
  }

  return nil
}
