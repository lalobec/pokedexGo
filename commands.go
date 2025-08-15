package main

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
