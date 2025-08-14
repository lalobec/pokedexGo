package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io"
  "log"
)


func CommandMap(c *config) {
  if c.Next == "" {
    c.Next = "https://pokeapi.co/api/v2/location-area/"
  }
  res, err := http.Get(c.Next)
  if err != nil {
    log.Fatal(err)
  }
  body, err := io.ReadAll(res.Body)
  res.Body.Close()

  var locationArea dataLoc 
  err2 := json.Unmarshal(body, &locationArea)
  if err2 != nil {
    fmt.Printf("Error in unmarshal ... %v \n", err2)
  }

  for _, loc := range locationArea.Results {
    fmt.Println(loc.Name)
  }

  c.Next = locationArea.Next
  c.Previous = locationArea.Previous
  //fmt.Printf("Testing the pointer to next ... %v \n", c.Next)

}

func CommandMapBack(c *config) {
  if c.Previous == "" {
    c.Previous = "https://pokeapi.co/api/v2/location-area/"
  }
  res, err := http.Get(c.Previous)
  if err != nil {
    log.Fatal(err)
  }
  body, err := io.ReadAll(res.Body)
  res.Body.Close()

  var locationArea dataLoc 
  err2 := json.Unmarshal(body, &locationArea)
  if err2 != nil {
    fmt.Printf("Error in unmarshal ... %v \n", err2)
  }

  for _, loc := range locationArea.Results {
    fmt.Println(loc.Name)
  }

  c.Next = locationArea.Next
  c.Previous = locationArea.Previous
  //fmt.Printf("Testing the pointer to next ... %v \n", c.Next)

}
