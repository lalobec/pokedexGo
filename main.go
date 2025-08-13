package main

import (
  "fmt"
  "bufio"
  "os"
  //"strings"
)

func main()  {
  scanner := bufio.NewScanner(os.Stdin)
  var input string
  for true {
    fmt.Print("Pokedex > ")
    if scanner.Scan() {
      input = scanner.Text()
    }
    firstWord := cleanInput(input)[0]
    fmt.Printf("Your command was: %s \n", firstWord)
  }
}


