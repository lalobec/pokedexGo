package main

import (
	"bufio"
	"fmt"
	"github.com/lalobec/pokedexGo/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocs      *string
	previousLocs  *string
	pokedex	map[string]pokeapi.PokemonType
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
				continue
			}

		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	var splitText []string
	lowerTextSlice := strings.Split(strings.ToLower(text), " ")
	for _, word := range lowerTextSlice {
		if word == "" {
			continue
		}
		trimmedWord := strings.TrimSpace(word)
		splitText = append(splitText, trimmedWord)
	}
	return splitText
}

/* //cleanInput could be done much more simple:
func cleanInput(text string) []string {
  output := strings.ToLower(text)
  words := strings.Fields(output) // this does all!
  return words
}
*/
