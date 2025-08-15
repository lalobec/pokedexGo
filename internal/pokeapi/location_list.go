package pokeapi

import (
	"encoding/json"
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

	// Here comes the request
  // First we create it
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

  locationsResp := DecLocations{}
  err = json.Unmarshal(dat, &locationsResp)
  if err != nil {
    return DecLocations{}, err
  }
  return locationsResp, nil
}
