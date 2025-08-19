package pokeapi

import (
	//"fmt"
	"testing"
  "time"
)

func TestFirstLocRequest(t *testing.T) {
  client := NewClient(5*time.Second, 5*time.Second)
  pokeList, err := client.PokemonList("https://pokeapi.co/api/v2/location-area/canalave-city-area") 
  
  if err != nil {
    t.Errorf("Something is wrong with the location request")
    return
  }
  if pokeList != nil {
    t.Errorf("This is what we got \n%v", pokeList)
    return 
  }
}
