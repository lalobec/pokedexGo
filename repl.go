package main

import (
  "strings"
  "fmt"
  "bufio"
  "os"
)

func startRepl() {
  scanner := bufio.NewScanner(os.Stdin)
  var input string
  for true {
    fmt.Print("Pokedex > ")
    if scanner.Scan() {
      input = scanner.Text()
      if len(input) == 0 {
        continue
      }
      firstWord := cleanInput(input)[0]
      fmt.Printf("Your command was: %s \n", firstWord)
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

// This function could be done much more simple:
/*
func cleanInput(text string) []string {
  output := strings.ToLower(text)
  words := strings.Fields(output) // this does all!
  return words
}
*/

