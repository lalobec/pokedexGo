package main

// This struct describes a cmd for our REPL
type cliCommand struct {
	name        string
	description string
	callback    func(c *config, args ...string) error
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
		"explore": {
			name:        "explore",
			description: "List the pokemons in the given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Use this command to throw a Pokeball to catch a Pokemon!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "See the details of the catched pokemon, you need to provide a name",
			callback:    commandInspect,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
