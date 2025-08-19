package main

import (
//	"errors"
  "fmt"
)

func commandExplore(cfg *config) error {
  url := "https://pokeapi.co/api/v2/location-area/canalave-city-area"
  pokeList, err := cfg.pokeapiClient.PokemonList(url)  
  if err != nil {
    return err
  }

  for _, pokemon := range pokeList {
    fmt.Println(pokemon)
  }

  return nil
}
