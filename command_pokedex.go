package main

import (
  "fmt"
	"errors"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("You have not catched any pokemon yet.\n")
	}	

	fmt.Println("---------------------------------")
	fmt.Println("------------ Pokedex ------------")
	fmt.Println("---------------------------------")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
