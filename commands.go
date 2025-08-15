package main

import (
	"fmt"
	"os"
)

// This struct describes a cmd for our REPL
type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show the next 20 locations of the map",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 locations of the map",
			callback:    commandMapb,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandExit(c *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Println("Usage:")
	cmdmap := getCommands()
	for _, cmd := range cmdmap {
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	return nil
}
