package main

import (
  "strings"
)

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

