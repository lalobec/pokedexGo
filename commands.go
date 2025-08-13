package main

import (
	"fmt"
	"os"
)

// This struct describes a cmd for our REPL
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Println("Usage:")
	cmdmap := getCommands()
	for _, cmd := range cmdmap {
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	return nil
}
