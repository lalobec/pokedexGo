package main

import (
  "fmt"
  "errors"
)


func commandMapf(c *config) error {

  // the passed c pointer to the config struct has a field pokeapiClient
  // which is a Client struct, so it has the method LocList... 
  locationsResp, err := c.pokeapiClient.LocList(c.nextLocs)
  if err != nil {
    return err
  }

  c.nextLocs = locationsResp.Next
  c.previousLocs = locationsResp.Previous

  for _, loc := range locationsResp.Results {
    fmt.Println(loc.Name)
  }
  return nil
}

func commandMapb(c *config) error {
  if c.previousLocs == nil {
    return errors.New("You are on the first page")
  }

  locationsResp, err := c.pokeapiClient.LocList(c.previousLocs)
  if err != nil {
    return err
  }

  c.nextLocs = locationsResp.Next
  c.previousLocs = locationsResp.Previous

  for _, loc := range locationsResp.Results {
    fmt.Println(loc.Name)
  }
  return nil
 }

