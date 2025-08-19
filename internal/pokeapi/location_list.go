package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Note that this is a method for the Client struct. So we can
// use it as clientName.LocList(url)
func (c *Client) LocList(pageURL *string) (DecLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL // In case we have another location url ??
	}

	if val, ok := c.cacheClient.Get(url); ok {

		locationsResp := DecLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return DecLocations{}, err
		}
    fmt.Println("---------------------------")
		fmt.Println("Locations were in the cache")
    fmt.Println("---------------------------")
		return locationsResp, nil

	} else {

		// Here comes the request. First we create it
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return DecLocations{}, err
		}

		// Then we do the request
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return DecLocations{}, err
		}
		defer resp.Body.Close()

		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return DecLocations{}, err
		}

		c.cacheClient.Add(url, dat)

		locationsResp := DecLocations{}
		err = json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return DecLocations{}, err
		}
    fmt.Println("------------------------------------")
		fmt.Println("Locations were requested to the web")
    fmt.Println("------------------------------------")
		return locationsResp, nil
	}
}
