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
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	for true {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input = scanner.Text()
			if len(input) == 0 {
				continue
			}

			commandName := cleanInput(input)[0]

			command, exists := getCommands()[commandName]
			if exists {
				if commandName == "explore" {
					err := command.callback(cfg, cleanInput(input)[1])
					if err != nil {
						fmt.Println(err)
						continue
					}

				} else {
					err := command.callback(cfg, "")
					if err != nil {
						fmt.Println(err)
						continue
					}
				}
			} else {
				fmt.Println("Unknown command")
				continue
			}
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
