package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonStats(pokemonName string) (PokemonType, error) {

	url := baseURL + "/pokemon/" + pokemonName

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonType{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonType{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonType{}, err
	}

	pokemonStats := PokemonType{}
	err = json.Unmarshal(dat, &pokemonStats)
	if err != nil {
		return PokemonType{}, err
	}

	fmt.Println("----------------------------------")
	fmt.Println("Pokemon data requested to the web")
	fmt.Println("----------------------------------")

	return pokemonStats, nil
}
