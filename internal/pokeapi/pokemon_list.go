package pokeapi

import (
	"encoding/json"
//	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonList(location string) ([]string, error) {
  
	url := baseURL + "/location-area/" + location
  var pokemonList []string

	// Create the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []string{}, err
	}

	// Do the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	locDetails := LocationDetails{}
	err = json.Unmarshal(dat, &locDetails)
	if err != nil {
		return []string{}, err
	}

  for _, poke := range locDetails.PokemonEncounters {
    pokemonList = append(pokemonList, poke.Pokemon.Name)
  }

  return pokemonList, nil
}
