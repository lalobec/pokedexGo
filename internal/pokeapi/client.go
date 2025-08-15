package pokeapi

// This is completely copied from the solution
// of boot.dev, I am trying to improve my code
// based on the solution. Instead of using a
// simple http.GET method, we do it with a client.

import (
  "net/http"
  "time"
)

type Client struct {
  httpClient http.Client
}

// If the API takes more time than timeout due to any
// reason, the request will be cancelled to prevent the 
// CLI from hanging indefinitely
func NewClient(timeout time.Duration) Client {
  return Client {
    httpClient: http.Client{
      Timeout: timeout,
    },
  }
}


