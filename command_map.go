package main

import (
  "fmt"
  "errors"
)


func commandMapf(cfg *config, s string) error {

  // the passed cfg pointer to the config struct has a field pokeapiClient
  // which is a Client struct, so it has the method LocList... 
  locationsResp, err := cfg.pokeapiClient.LocList(cfg.nextLocs)
  if err != nil {
    return err
  }

  cfg.nextLocs = locationsResp.Next
  cfg.previousLocs = locationsResp.Previous

  for _, loc := range locationsResp.Results {
    fmt.Println(loc.Name)
  }
  return nil
}

func commandMapb(cfg *config, s string) error {
  if cfg.previousLocs == nil {
    return errors.New("You are on the first page")
  }

  locationsResp, err := cfg.pokeapiClient.LocList(cfg.previousLocs)
  if err != nil {
    return err
  }

  cfg.nextLocs = locationsResp.Next
  cfg.previousLocs = locationsResp.Previous

  for _, loc := range locationsResp.Results {
    fmt.Println(loc.Name)
  }
  return nil
 }

