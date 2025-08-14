package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
  c_real := config{
    Next: "",
    Previous: "",
  }
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
				err := command.callback(&c_real)
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
