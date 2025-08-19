package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocDetails(url string) (LocationDetails, error) {

	// Create the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	// Do the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	locDetails := LocationDetails{}
	err = json.Unmarshal(dat, &locDetails)
	if err != nil {
		return LocationDetails{}, err
	}

  fmt.Println("It is time for some testing")
  return locDetails, nil
}
