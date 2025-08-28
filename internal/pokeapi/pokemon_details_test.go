package pokeapi

import (
	"testing"
	"time"
)

func TestPikachu(t *testing.T) {
  client := NewClient(5*time.Second, 5*time.Second)
  pokemon, err := client.PokemonStats("pikachu") 
	
	expected := "pikachu"

	if err != nil {
		t.Errorf("Something bad happened requesting Pikachu")
		return
	}

	if pokemon.Name != expected {
		t.Errorf("Fail. Expecting %s, We got %s instead.", expected, pokemon.Name)
	}
} 

func TestPikachuExp(t *testing.T) {
  client := NewClient(5*time.Second, 5*time.Second)
  pokemon, err := client.PokemonStats("pikachu") 
	
	expected := 112 

	if err != nil {
		t.Errorf("Something bad happened requesting Pikachu data")
		return
	}

	if pokemon.BaseExperience != expected {
		t.Errorf("Fail. Expecting %d, We got %d instead.", expected, pokemon.BaseExperience)
	}
} 
