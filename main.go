package main

import (
  "fmt"
  "strings"
)

func main()  {
  fmt.Println("Hello, World!")
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

