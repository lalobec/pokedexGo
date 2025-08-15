package main

import (
  "time"
	"github.com/lalobec/pokedexGo/internal/pokeapi"
)

func main() {
  pokeClient := pokeapi.NewClient(5*time.Second)

  // Create a config struct and store it in c, that will
  // hold the pokeClient.
  c := &config{
    pokeapiClient: pokeClient,
  }

  // Pass the c config to the REPL so that it has
  // access to the pokeapi client.
	startRepl(c)
}
