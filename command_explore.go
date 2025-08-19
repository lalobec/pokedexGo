package main

import (
//	"errors"
  "fmt"
)

func commandExplore(cfg *config, location string) error {
  pokeList, err := cfg.pokeapiClient.PokemonList(location)  
  if err != nil {
    return err
  }

  for _, pokemon := range pokeList {
    fmt.Println(pokemon)
  }

  return nil
}
