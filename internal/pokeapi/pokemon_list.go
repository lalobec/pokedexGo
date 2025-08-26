package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonList(locationName string) ([]string, error) {

	url := baseURL + "/location-area/" + locationName
	var pokemonList []string

	// Checking if data is in cache
	if val, ok := c.cacheClient.Get(url); ok {
		locDetails := LocationDetails{}
		err := json.Unmarshal(val, &locDetails)
		if err != nil {
			return []string{}, err
		}

		for _, poke := range locDetails.PokemonEncounters {
			pokemonList = append(pokemonList, poke.Pokemon.Name)
		}

		fmt.Println("---------------------------")
		fmt.Println("Pokemons were in the cache")
		fmt.Println("---------------------------")

		return pokemonList, nil

	} else {

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

    c.cacheClient.Add(url, dat)

		locDetails := LocationDetails{}
		err = json.Unmarshal(dat, &locDetails)
		if err != nil {
			return []string{}, err
		}

		for _, poke := range locDetails.PokemonEncounters {
			pokemonList = append(pokemonList, poke.Pokemon.Name)
		}
		
		fmt.Println("------------------------------")
		fmt.Println("Pokemons requested to the web")
		fmt.Println("------------------------------")

		return pokemonList, nil
	}
}
