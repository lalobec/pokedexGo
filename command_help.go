package main

import (
  "fmt"
)

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Println("Usage:")
	cmdmap := getCommands()
	for _, cmd := range cmdmap {
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	return nil
}
