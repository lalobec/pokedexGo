package main

import (
  "fmt"
  "os"
)

func commandExit(c *config, s string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
