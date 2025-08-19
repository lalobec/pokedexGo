package pokeapi

import (
	//"fmt"
	"testing"
  "time"
)

func TestFirstLocRequest(t *testing.T) {
  client := NewClient(5*time.Second, 5*time.Second)
  loc_one, err := client.LocDetails("https://pokeapi.co/api/v2/location-area/canalave-city-area") 
  
  if err != nil {
    t.Errorf("Something is wrong with the location request")
    return
  }
  if loc_one.Location.Name != "canalave-city" {
    t.Errorf("Expecting 'canalave-city', Got %s", loc_one.Location.Name)
    return 
  }
}
