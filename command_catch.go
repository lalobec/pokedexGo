package main

import (
  "fmt"
	"errors"
	"math/rand"
	"time"
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
	time.Sleep(500 * time.Millisecond)
	fmt.Print(" . ")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(" . ")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(" . \n")

	pokeProbability := float64(pokemon.BaseExperience) * 0.95 / 200
	pokeThrow := float64(rand.Intn(100))/100

	if pokeThrow < pokeProbability {
		fmt.Printf("%s ran away!, prob: %f, catching: %f\n", pokemon.Name, pokeProbability, pokeThrow)
		return nil
	}

	fmt.Printf("You catched %s!, prob: %f, catching: %f\n", pokemon.Name, pokeProbability, pokeThrow)
	cfg.pokedex[pokemon.Name] = pokemon
	
	/*
	fmt.Printf("Your pokedex has")
	for _, pokemon := range cfg.pokedex {
		fmt.Println(pokemon.Name)
	}
	*/

  return nil
}
